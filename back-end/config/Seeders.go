package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_"github.com/mattn/go-sqlite3"
)

func Seeders(databaseFile string, seederFile string) error {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Read the seed SQL file
	seeder, err := ioutil.ReadFile(seederFile)
	if err != nil {
		return fmt.Errorf("failed to read seeder file: %v", err)
	}

	// Execute the SQL commands in the seeder file
	_, err = db.Exec(string(seeder))
	if err != nil {
		return fmt.Errorf("failed to execute seeder: %v", err)
	}

	return nil
}
