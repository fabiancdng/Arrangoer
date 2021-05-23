package database

import (
	"github.com/fabiancdng/Arrangoer/internal/models"
)

// Definiert, welche Funktionen eine Datenbank-Middleware vorweisen muss
type Middleware interface {
	// Prüfen, ob die Datenbank-Datei
	Prepare() error
	// Anmeldung für den Wettbewerb in der Datenbank speichern
	SaveApplication(*models.Application) error
}
