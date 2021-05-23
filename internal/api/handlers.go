package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// +++++++ AUTHENTICATION HANDLERS +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// Leite an die Oauth2 Authorization Seite weiter
func (api *API) auth(ctx *fiber.Ctx) error {
	return ctx.Redirect(api.discordAuth.AuthCodeURL(api.state), 307)
}

func (api *API) authCallback(ctx *fiber.Ctx) error {
	callbackRequest := new(CallbackRequest)

	err := ctx.QueryParser(callbackRequest)
	if err != nil {
		return fiber.NewError(400, "invalid request body")
	}

	if callbackRequest.Sate != api.state {
		return fiber.NewError(400, "state doesn't match")
	}

	// Den Code für einen Access-Token eintauschen
	token, err := api.discordAuth.Exchange(context.Background(), callbackRequest.Code)

	if err != nil {
		return fiber.NewError(500, "an error occured at code/token exchange")
	}

	sess, err := api.store.Get(ctx)
	if err != nil {
		return fiber.NewError(500, "error processing session")
	}

	defer sess.Save()

	// Discord Access- & Refresh-Token in serverseitiger Session speichern
	sess.Set("dc_access_token", token.AccessToken)
	sess.Set("dc_refresh_token", token.AccessToken)

	return ctx.Redirect("http://localhost:3000")
}

// Daten von der Discord OAuth2 API abrufen (wie Nutzerinfos oder Guildinfos)
func (api *API) authGetFromEndpoint(ctx *fiber.Ctx) error {
	accessToken, refreshToken, err := authorize(ctx, api.store)
	if accessToken == "" || refreshToken == "" || err != nil {
		return fiber.NewError(401)
	}

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
	res, err := api.discordAuth.Client(context.Background(), token).Get(endpoint)

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
			if guild.ID == api.config.Discord.ServerID {
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
			"invite_link":    api.config.Discord.InviteLink,
		})

	default:
		var discordUser DiscordUser
		json.Unmarshal(body, &discordUser)

		return ctx.JSON(discordUser)
	}

}

func (api *API) authLogout(ctx *fiber.Ctx) error {
	sess, err := api.store.Get(ctx)
	if err != nil || sess == nil {
		return fiber.NewError(500, "error processing session")
	}

	if err := sess.Destroy(); err != nil {
		return fiber.NewError(500, "error deleting session")
	}

	return ctx.SendStatus(200)
}

// +++++++ APPLICATION HANDLERS +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func (api *API) applicationSubmit(ctx *fiber.Ctx) error {
	accessToken, refreshToken, err := authorize(ctx, api.store)
	if accessToken == "" || refreshToken == "" || err != nil {
		return fiber.NewError(401)
	}

	application := new(Application)

	err = ctx.BodyParser(application)
	if err != nil {
		return fiber.NewError(400)
	}

	log.Println(application)

	return ctx.SendStatus(200)
}
