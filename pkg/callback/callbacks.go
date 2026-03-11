package callback

import (
	"fmt"
	"log"
	"os"

	pokeapi "github.com/humbl3LilaS/pokedexcli/api"
)

func callbackHelp() error {
	fmt.Println("Welcome To The Pokedex Cli Menu....")
	fmt.Println("Available Command Options:")
	availCmds := GetCommands()

	for _, cmd := range availCmds {
		fmt.Printf("- %s: %s\n", cmd.Name, cmd.Desctiption)
	}

	return nil
}

func callbackExit() error {
	fmt.Println("Exiting Pokedex CLI....")
	os.Exit(0)
	return nil
}


func callbackMap() error {
	apiClient := pokeapi.NewClient()

	resp, err := apiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Location Areas: ")

	for _, loc := range resp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}
