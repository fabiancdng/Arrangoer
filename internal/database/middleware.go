package database

import (
	"github.com/fabiancdng/Arrangoer/internal/models"
)

// Definiert, welche Funktionen eine Datenbank-Middleware vorweisen muss
type Middleware interface {
	Open() error
	Close() error
	// Anmeldung für den Wettbewerb in der Datenbank speichern
	SaveApplication(application *models.ApplicationRequest) error
	// Alle Anmeldungen einsehen
	GetApplications() ([]models.Application, error)
}
