package callback

import (
	"fmt"
	"os"
)

func callbackHelp() {
	fmt.Println("Welcome To The Pokedex Cli Menu....")
	fmt.Println("Available Command Options:")
	availCmds := GetCommands()

	for _, cmd := range availCmds {
		fmt.Printf("- %s: %s\n", cmd.Name, cmd.Desctiption)
	}
}

func callbackExit() {
	fmt.Println("Exiting Pokedex CLI....")
	os.Exit(0)
}

