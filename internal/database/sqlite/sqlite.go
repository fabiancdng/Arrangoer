package sqlite

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"github.com/fabiancdng/Arrangoer/internal/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

// Datenbank Middleware für SQLite
type SQLite struct {
	db *sql.DB
}

// Öffnet die Datenbank und erstellt (falls noch nicht vorhanden) alle nötigen Tabellen
func (sqlite *SQLite) Open() error {
	var err error

	if _, err := os.Stat("./db.db"); os.IsNotExist(err) {
		err := ioutil.WriteFile("./db.db", []byte(""), 0755)
		if err != nil {
			return err
		}
		log.Println("Datenbank-Datei wurde erstellt.")
	}

	sqlite.db, err = sql.Open("sqlite3", "./db.db")
	if err != nil {
		return err
	}

	_, err = sqlite.db.Exec("CREATE TABLE IF NOT EXISTS `teams` (`id` INTEGER PRIMARY KEY, `name` VARCHAR(50), `approved` INT)")
	if err != nil {
		return err
	}

	_, err = sqlite.db.Exec("CREATE TABLE IF NOT EXISTS `applications` (`id` INTEGER PRIMARY KEY, `name` VARCHAR(100), `email` VARCHAR(200), `team` INT, `user_id` VARCHAR(100), `accepted` INT, FOREIGN KEY (`team`) REFERENCES `teams`(`id`))")
	if err != nil {
		return err
	}

	return nil
}

// Schließt die Verbindung zur Datenbank
func (sqlite *SQLite) Close() error {
	if err := sqlite.db.Close(); err != nil {
		return err
	}
	return nil
}

// Speichert eine Anmeldung für den Wettbewerb in der Datenbank
// Ebenso wie das Team (bzw. erstellt es ggf.)
func (sqlite *SQLite) SaveApplication(application *models.Application) error {
	// Prüfen, ob der Nutzer sich bereits angemeldet hat
	id := 0
	err := sqlite.db.QueryRow("SELECT `id` FROM `applications` WHERE `name`=?", application.Team).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return fiber.NewError(500)
		}
		// Nutzer ist noch nicht angemeldet
	} else {
		// Nutzer ist bereits angemeldet
		return fiber.NewError(302)
	}

	// Prüfen, ob das eingegebene Team bereits existiert
	err = sqlite.db.QueryRow("SELECT `id` FROM `teams` WHERE `name`=?", application.Team).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return fiber.NewError(500)
		}

		// Das Team existiert (in der DB) noch nicht und muss daher erstellt werden
		_, err = sqlite.db.Exec("INSERT INTO `applications` VALUES (NULL, ?, ?, ?, ?, 0)", application.Name, application.Email, application.Team, application.UserID)
		if err != nil {
			return fiber.NewError(500)
		}

	} else {
		// Das Team existiert (in der DB) bereits
		return fiber.NewError(302)
	}

	// Die Anmeldung in der Datenbank speichern
	// Das Team wird über eine Team-ID (die auf eine andere Tabelle weist) referenziert
	_, err = sqlite.db.Exec("INSERT INTO `applications` VALUES (NULL, ?, ?, ?, ?, 0)", application.Name, application.Email, application.Team, application.UserID)
	if err != nil {
		return fiber.NewError(500)
	}

	log.Printf("%s hat sich für den Wettbewerb angemeldet.", application.Name)

	return nil
}

// Gibt eine Liste mit allen Anmeldungen (aus der Datenbank) zurück
func (sqlite *SQLite) GetApplications() ([]models.Application, error) {
	rows, err := sqlite.db.Query("SELECT * FROM `applications`")
	if err != nil {
		return nil, fiber.NewError(500)
	}

	currentApplication := new(models.Application)
	applications := []models.Application{}

	for rows.Next() {
		rows.Scan(&currentApplication.ID, &currentApplication.Name, &currentApplication.Email, &currentApplication.Team, &currentApplication.UserID, &currentApplication.Accepted)
		applications = append(applications, *currentApplication)
	}
	log.Println(applications)
	return applications, nil
}
