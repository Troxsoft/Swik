package main

import (
	"fmt"
	"os"

	"github.com/Troxsoft/Swik/cmd"
	"github.com/Troxsoft/Swik/pkg"
)

func main() {
	args := os.Args
	if len(args) == 3 {
		if args[1] == "new" {

			err := pkg.CreateDB(args[2])
			if err != nil {
				fmt.Printf("Error: %v ❌\n", err)
			} else {
				fmt.Printf("Database created successfully: %s ✔", args[2])
			}
		} else {
			fmt.Printf("Command invalid: %v\n", args[1:])
		}
	} else if len(args) == 2 {
		db, err := pkg.NewDBFromFile(args[1])
		if err != nil {
			fmt.Printf("Error: %v ❌\n", err)
		}
		db.AddJSFunctions()
		cli := cmd.NewCLI(db)
		cli.Run()

	} else {
		fmt.Printf("Comand invalid: %v\n", args[1:])
	}

}
