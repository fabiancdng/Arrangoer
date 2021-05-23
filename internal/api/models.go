package api

// Discord User in OAuth2 API
type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
}

// Discord Guild in OAuth2 API
type DiscordGuild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Permissions string `json:"permissions_new"`
}

// Eine Anfrage auf die Seite, an die Discord einen nach der Autorisierung weiterleitet
type CallbackRequest struct {
	Sate string `query:"state"`
	Code string `query:"code"`
}

// Daten für eine Anmeldung zum Wettbewerb, die über das Formular abgeschickt werden
type Application struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Team  string `json:"team"`
}
