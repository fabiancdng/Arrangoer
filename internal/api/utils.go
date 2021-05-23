package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Funktion, die prüft, ob der Nutzer (gültig) mit Discord angemeldet ist
func authorize(ctx *fiber.Ctx, store *session.Store) (string, string, error) {
	sess, err := store.Get(ctx)
	if err != nil || sess == nil {
		return "", "", fiber.NewError(500, "error processing session")
	}

	defer sess.Save()

	sessionAccessToken := sess.Get("dc_access_token")
	if sessionAccessToken == nil {
		return "", "", fiber.NewError(401)
	}

	accessToken := sessionAccessToken.(string)

	sessionRefreshToken := sess.Get("dc_refresh_token")
	if sessionAccessToken == nil {
		return accessToken, "", fiber.NewError(401)
	}

	refreshToken := sessionRefreshToken.(string)

	return accessToken, refreshToken, nil
}
