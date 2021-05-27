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
	// Gibt eine bestimmte Anmeldung zurück
	GetApplication(applicationID int) (*models.Application, error)
	// Gibt alle Member sowie den Namen eines Teams zurück
	GetTeam(teamID int) (*models.Team, error)
	// Alle Anmeldungen zurückgeben
	GetApplications() ([]models.Application, error)
	// Alle Teams zurückgeben
	GetTeams() ([]models.Team, error)
	// Eine Anmeldung in der Datenbank als 'akzeptiert' markieren
	AcceptApplication(applicationID int, applicantName string) error
	// Ein Team in der Datenbank als 'akzeptiert' markieren
	ApproveTeam(teamID int, teamName string) error
	// Eine Anmeldung aus der Datenbank löschen, da sie abgelehnt wurde
	DeclineApplication(applicationID int) error
	// Ein Team aus der Datenbank löschen, da es abgelehnt wurde
	DeclineTeam(teamID int) error
}
