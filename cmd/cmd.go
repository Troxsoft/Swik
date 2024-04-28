package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Troxsoft/Swik/pkg"
)

type CLI struct {
	DB      *pkg.DB
	running bool
}

func (cli *CLI) Run() {
	cli.running = true

	for cli.running {

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.HasPrefix(text, "@save") {
			err := cli.DB.Save()
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
			} else {
				fmt.Printf("Database saved successfully ✔\n")
			}

		} else if strings.HasPrefix(text, "@close") {
			return

		} else {
			v, err := cli.DB.JS().Run(text)
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
			} else {

				fmt.Println(v.String())
			}
		}
	}
}
func NewCLI(db *pkg.DB) *CLI {
	return &CLI{
		DB:      db,
		running: false,
	}
}
