package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

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
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}

	app := fiber.New()
	store := session.New()

	app.Get("/api/auth", func(ctx *fiber.Ctx) error {
		// Leite an die Oauth2 Authorization Seite weiter
		return ctx.Redirect(discordAuth.AuthCodeURL(state), 307)
	})

	app.Get("/api/auth/get", func(ctx *fiber.Ctx) error {
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

		// Den Access-Token benutzen, um Daten des Benutzers abzurufen
		res, err := discordAuth.Client(context.Background(), token).Get("https://discordapp.com/api/users/@me")

		if err != nil || res.StatusCode != 200 {
			return fiber.NewError(500, "couldn't use the access token")
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return fiber.NewError(500, "an error occured while attempting to parse request body")
		}

		var discordUser DiscordUser
		json.Unmarshal(body, &discordUser)

		return ctx.JSON(discordUser)
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
