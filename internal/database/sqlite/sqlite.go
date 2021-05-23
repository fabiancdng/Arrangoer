package sqlite

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Datenbank Middleware für SQLite
type SQLite struct {
	db *sql.DB
}

func (sqlite *SQLite) Open() error {
	var err error

	if _, err := os.Stat("./database.db"); os.IsNotExist(err) {
		err := ioutil.WriteFile("./database.db", []byte(""), 0755)
		if err != nil {
			return err
		}
		log.Println("Datenbank-Datei wurde erstellt.")
	}

	sqlite.db, err = sql.Open("sqlite3", "./database.db")
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
