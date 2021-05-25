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
	// Eine Anmeldung in der Datenbank als 'akzeptiert' markieren
	AcceptApplication(applicationID int) error
	// Ein Team in der Datenbank als 'akzeptiert' markieren
	ApproveTeam(teamID int) error
	// Eine Anmeldung aus der Datenbank löschen, da sie abgelehnt wurde
	DeclineApplication(applicationID int) error
	// Ein Team aus der Datenbank löschen, da es abgelehnt wurde
	DeclineTeam(teamID int) error
}
