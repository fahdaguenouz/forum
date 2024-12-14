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
	if cmd == "--migrate" || cmd == "-m" {
		err := Migrate("./back-end/database/db.sql", "./back-end/database/database.db")
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Database migration completed successfully.")
		return
	} else if cmd == "--seed" {
		err := Seeders("./back-end/database/database.db", "./back-end/database/seeder.sql")
		if err != nil {
			log.Fatalf("Seeder failed: %v", err)
		}
		fmt.Println("Seeder completed successfully.")
	} else {
		fmt.Println("Invalid command try : --migrate or --seed")
		return
	}
}
