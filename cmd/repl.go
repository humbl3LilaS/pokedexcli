package cmd

import (
	"bufio"
	"fmt"
	"os"

	util "github.com/humbl3LilaS/pokedexcli/helper"
)

func StartRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >")
		scanner.Scan()
		input := scanner.Text();

		clean := util.CleanInput(input)

		if len(clean) == 0 {
			continue
		}

		command := clean[0]

		switch command {
		case "help":
			fmt.Println("Welcome To The Pokedex Cli Menu....")
			fmt.Println("Available Command Options:")
			fmt.Println("- help: Print out the available command options")
			fmt.Println("- exit: Exit from Pokedex Cli")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Command. Use `help` to check for available options.")
		}

	}
}


