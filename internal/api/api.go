package api

import (
	"log"

	"github.com/fabiancdng/Arrangoer/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
)

type API struct {
	// Fiber App (webserver, fiber middlewares, etc.)
	app *fiber.App
	// Store für sessions
	store *session.Store
	// Geparste config.yml
	config *config.Config
	// Config für Discord OAuth2 Middleware
	discordAuth *oauth2.Config
	// Zufälliger String, der vom Login geschickt und im Callback validiert wird
	state string
}

func NewAPI(config *config.Config) (*API, error) {
	var state string = "v6uhSq6eWsnyAp"

	discordAuth := &oauth2.Config{
		RedirectURL:  "http://localhost:5000/api/auth/callback",
		ClientID:     config.Discord.ClientID,
		ClientSecret: config.Discord.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	app := fiber.New()
	store := session.New()

	api := &API{
		app:         app,
		store:       store,
		config:      config,
		discordAuth: discordAuth,
		state:       state,
	}

	api.registerHandlers()

	return api, nil
}

func (api *API) registerHandlers() {
	// Hauptgruppe für alle API Endpoints
	// Routes für /api/*
	apiGroup := api.app.Group("/api")

	// Untergruppe für Authentication Endpoints
	// Routes für /api/auth/*
	apiAuthGroup := apiGroup.Group("/auth")
	apiAuthGroup.Get("/", api.auth)
	apiAuthGroup.Get("/callback", api.authCallback)
	apiAuthGroup.Get("/get/:endpoint", api.authGetFromEndpoint)
	apiAuthGroup.Get("/logout", api.authLogout)

	// Untergruppe für Anmeldungs Endpoints
	// Routes für /api/application/*
	apiApplicationGroup := apiGroup.Group("/application")
	apiApplicationGroup.Post("/submit", api.applicationSubmit)
}

func (api *API) RunAPI(apiChannel chan string) {
	log.Println("API ist bereit!")
	api.app.Listen(api.config.API.AddressAndPort)
}
