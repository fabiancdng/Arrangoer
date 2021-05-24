package models

// Daten für eine Anmeldung zum Wettbewerb, die über das Formular abgeschickt werden
type Application struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	UserID   string `json:"user_id"`
	Accepted int    `json:"accepted"`
}
