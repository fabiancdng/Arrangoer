package database

// Definiert, welche Funktionen eine Datenbank Middleware vorweisen muss
type Middleware interface {
	// Datenbankverbindung aufbauen und Tabellen vorbereiten
	Open() error
}
