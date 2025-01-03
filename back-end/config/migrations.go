package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_"github.com/mattn/go-sqlite3"
)

func Migrate(schemaFile string, databaseFile string) error {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	// Read the schema SQL file
	schema, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Execute the SQL commands in the schema file
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}

	return nil

}
