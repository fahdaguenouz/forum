package config

import (
	"fmt"
	"log"
	
)

func Config(args []string) {
	if len(args) > 2 {
		fmt.Println("Invalid number of arguments")
		return
	}
	cmd := args[1]
	fmt.Println(cmd)
	if cmd == "--migrate" {
		err := Migrate("db.sql", "./database.sqlite")
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Database migration completed successfully.")
		return
	} else if cmd == "--seed" {
		Seeders()
	} else {
		fmt.Println("Invalid command try : --migrate or --seed")
		return
	}
}
