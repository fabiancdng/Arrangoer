package sqlite

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"github.com/fabiancdng/Arrangoer/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// Datenbank Middleware für SQLite
type SQLite struct {
	db *sql.DB
}

func (sqlite *SQLite) Prepare() error {
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

	_, err = sqlite.db.Exec("CREATE TABLE IF NOT EXISTS `applications` (`id` INTEGER PRIMARY KEY, `name` VARCHAR(100), `email` VARCHAR(200), `team` VARCHAR(100))")
	if err != nil {
		return err
	}

	sqlite.db.Close()

	return nil
}

func (sqlite *SQLite) SaveApplication(application *models.Application) error {
	_, err := sqlite.db.Exec("INSERT INTO `applications` VALUES (NULL, %s, %s, %s)", application.Name, application.Email, application.Team)
	if err != nil {
		return err
	}

	return nil
}
