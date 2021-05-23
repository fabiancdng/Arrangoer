package database

import "github.com/fabiancdng/Arrangoer/internal/api"

// Definiert, welche Funktionen eine Datenbank Middleware vorweisen muss
type DatabaseMiddleware interface {
	Open() error
	Close()

	SaveApplication(application *api.Application) error
}
