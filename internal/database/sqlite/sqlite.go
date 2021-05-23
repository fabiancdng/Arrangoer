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

	_, err = sqlite.db.Exec("CREATE TABLE IF NOT EXISTS `applications` (`id` INTEGER PRIMARY KEY, `name` VARCHAR(100), `email` VARCHAR(200), `team` VARCHAR(100), `user_id` VARCHAR(100))")
	if err != nil {
		return err
	}

	return nil
}

func (sqlite *SQLite) Close() error {
	if err := sqlite.db.Close(); err != nil {
		return err
	}
	return nil
}

func (sqlite *SQLite) SaveApplication(application *models.Application) error {
	id := 0
	err := sqlite.db.QueryRow("SELECT `id` FROM `applications` WHERE `user_id`='?'", application.UserID).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return fiber.NewError(500)
		}

		return fiber.NewError(302)
	}

	log.Println(id, application.UserID)

	_, err = sqlite.db.Exec("INSERT INTO `applications` VALUES (NULL, ?, ?, ?, ?)", application.Name, application.Email, application.Team, application.UserID)
	if err != nil {
		return fiber.NewError(500)
	}

	return nil
}
