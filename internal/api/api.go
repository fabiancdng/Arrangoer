package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/fabiancdng/Arrangoer/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
}

type DiscordGuild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Permissions string `json:"permissions_new"`
}

type CallbackRequest struct {
	Sate string `query:"state"`
	Code string `query:"code"`
}

func Run(apiChannel chan string) {
	config, err := config.ParseConfig("./config/config.json")
	if err != nil {
		log.Panic(err)
	}

	// Zufälliger String, der von Login geschickt und im Callback validiert wird
	var state string = "v6uhSq6eWsnyAp"

	discordAuth := &oauth2.Config{
		RedirectURL:  "http://localhost:5000/api/auth/callback",
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	app := fiber.New()
	store := session.New()

	app.Get("/api/auth", func(ctx *fiber.Ctx) error {
		// Leite an die Oauth2 Authorization Seite weiter
		return ctx.Redirect(discordAuth.AuthCodeURL(state), 307)
	})

	app.Get("/api/auth/get/:endpoint", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil || sess == nil {
			return fiber.NewError(500, "error processing session")
		}

		defer sess.Save()

		sessionAccessToken := sess.Get("dc_access_token")
		if sessionAccessToken == nil {
			return fiber.NewError(401)
		}

		accessToken := sessionAccessToken.(string)

		sessionRefreshToken := sess.Get("dc_refresh_token")
		if sessionAccessToken == nil {
			return fiber.NewError(401)
		}

		refreshToken := sessionRefreshToken.(string)

		token := &oauth2.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		var endpoint string

		switch ctx.Params("endpoint") {
		case "guild":
			endpoint = "https://discordapp.com/api/users/@me/guilds"
		default:
			endpoint = "https://discordapp.com/api/users/@me"
		}

		// Den Access-Token benutzen, um Daten des Benutzers abzurufen
		res, err := discordAuth.Client(context.Background(), token).Get(endpoint)

		if err != nil || res.StatusCode != 200 {
			return fiber.NewError(500, "couldn't use the access token")
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return fiber.NewError(500, "an error occured while attempting to parse request body")
		}

		switch ctx.Params("endpoint") {
		case "guild":
			var discordGuilds []*DiscordGuild
			json.Unmarshal(body, &discordGuilds)

			var isUserMemberOfGuild bool = false
			var isUserAdminOfGuild bool = false

			for _, guild := range discordGuilds {
				// Prüfen, ob der Nutzer bereits auf dem Server ist
				if guild.ID == config.ServerID {
					isUserMemberOfGuild = true
					// Prüfen, ob der Nutzer Admin-Rechte auf dem Server hat
					permissions, _ := strconv.Atoi(guild.Permissions)
					if permissions&8 > 0 {
						isUserAdminOfGuild = true
					}

					break
				}
			}

			return ctx.JSON(fiber.Map{
				"user_is_member": isUserMemberOfGuild,
				"user_is_admin":  isUserAdminOfGuild,
			})

		default:
			var discordUser DiscordUser
			json.Unmarshal(body, &discordUser)

			return ctx.JSON(discordUser)
		}

	})

	app.Get("/api/auth/callback", func(ctx *fiber.Ctx) error {
		callbackRequest := new(CallbackRequest)

		err := ctx.QueryParser(callbackRequest)
		if err != nil {
			return fiber.NewError(400, "invalid request body")
		}

		if callbackRequest.Sate != state {
			return fiber.NewError(400, "state doesn't match")
		}

		// Den Code für einen Access-Token eintauschen
		token, err := discordAuth.Exchange(context.Background(), callbackRequest.Code)

		if err != nil {
			return fiber.NewError(500, "an error occured at code/token exchange")
		}

		sess, err := store.Get(ctx)
		if err != nil {
			return fiber.NewError(500, "error processing session")
		}

		defer sess.Save()

		// Discord Access- & Refresh-Token in serverseitiger Session speichern
		sess.Set("dc_access_token", token.AccessToken)
		sess.Set("dc_refresh_token", token.AccessToken)

		return ctx.Redirect("http://localhost:3000")
	})

	log.Println("API ist bereit!")

	app.Listen(":5000")
}
