package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Troxsoft/Swik/pkg"
	"github.com/fatih/color"
)

type CLI struct {
	DB      *pkg.DB
	running bool
}

func (cli *CLI) Run() {
	cli.running = true

	for cli.running {

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf(color.GreenString(">>"))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.HasPrefix(text, "@save") {
			err := cli.DB.Save()
			if err != nil {
				fmt.Printf("Error: %#v\n", err)
			} else {
				fmt.Printf("Database saved successfully ✔\n")
			}

		} else if strings.HasPrefix(text, "@exit") {
			return

		} else if strings.HasPrefix(text, "@cls") {
			fmt.Print("\033[H\033[2J")

		} else if strings.HasPrefix(text, "@js") {
			var codejs string
			for {
				text, _ = reader.ReadString('\n')
				if strings.TrimSpace(text) != "@js-end" {

					codejs += "\n" + text
				} else {
					break
				}
			}
			v, err := cli.DB.JS().RunString(codejs)

			if err != nil {
				fmt.Printf("Error: %#v\n", err.Error())
			} else {

				fmt.Printf("%s\n", color.YellowString(v.String()))

			}

		} else {

			v, err := cli.DB.JS().RunString(text)

			if err != nil {
				fmt.Printf("Error: %#v\n", err.Error())
			} else {

				fmt.Printf("%s\n", color.YellowString(v.String()))

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
