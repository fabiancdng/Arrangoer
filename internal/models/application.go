package models

// Daten für eine Anmeldung zum Wettbewerb, die über das Formular abgeschickt werden
type Application struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Team   string `json:"team"`
	UserID string
}
