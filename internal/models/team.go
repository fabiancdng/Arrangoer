package models

// Daten für ein Team
type Team struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Approved int    `json:"approved"`
}
