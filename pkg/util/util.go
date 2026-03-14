package util

import (
	"fmt"
	"strings"

	"github.com/humbl3LilaS/pokedexcli/pkg/callback"
	customTypes "github.com/humbl3LilaS/pokedexcli/pkg/types"
)

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words :=strings.Fields(lowered)
	return words
}

func callbackHelp(arg customTypes.CmdArg) error {
	fmt.Println("Welcome To The Pokedex Cli Menu....")
	fmt.Println("Available Command Options:")
	availCmds := GetCommands()

	for _, cmd := range availCmds {
		fmt.Printf("- %s: %s\n", cmd.Name, cmd.Desctiption)
	}

	return nil
}


func GetCommands() map[string]customTypes.CliCommand {
	return map[string]customTypes.CliCommand{
		"help": {
			Name: "help",
			Desctiption: "Print the help menu of Pokedex CLI.",
			CallBack: callbackHelp,
			MaxArg: 0,
		},
		"exit": {
			Name: "exit",
			Desctiption: "Turnoff the Pokedex CLI.",
			CallBack: callback.CallbackExit,
			MaxArg: 0,
		},
		"map": {
			Name: "map",
			Desctiption: "Output the Next Page of LocationAreas",
			CallBack: callback.CallbackMap,
			MaxArg: 0,
		},
		"mapb" : {
			Name: "mapb",
			Desctiption: "Output the Previous Page of LocationAreas",
			CallBack:callback.CallbackMapPrevious,
			MaxArg: 0,
		},
		"explore": {
			Name: "explore {locationID}",
			Desctiption: "Output the detail of the LocationArea",
			CallBack: callback.CallbackExplore,
			MaxArg: 1,
		},
		"catch": {
			Name:  "catch {pokemonName}",
			Desctiption: "Catch the pokemon",
			CallBack: callback.CallbackCatch,
			MaxArg: 1,
		},
	}
}

