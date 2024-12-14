package config

import "fmt"

func Config(args []string) {
	if len(args) > 2 {
		fmt.Println("Invalid number of arguments")
        return
	}
	cmd:=args[1]
	fmt.Println(cmd)
	if cmd == "--migrate"{
		Migrate()
	}else if cmd =="--seed"{
		Seeders()
	}else{
		fmt.Println("Invalid command try : --migrate or --seed")
        return
	}
}
