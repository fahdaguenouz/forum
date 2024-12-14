package main

import (
	"Forum/back-end/config"
	"Forum/back-end/routes"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		routes.Router()
	} else if len(args) > 1 {
		config.Config(args)
	}
}
