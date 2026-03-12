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

func callbackHelp(conf *customTypes.AppConfig) error {
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
		},
		"exit": {
			Name: "exit",
			Desctiption: "Turnoff the Pokedex CLI.",
			CallBack: callback.CallbackExit,
		},
		"map": {
			Name: "map",
			Desctiption: "Output the Next Page of LocationAreas",
			CallBack: callback.CallbackMap,
		},
		"mapb" : {
			Name: "mapb",
			Desctiption: "Output the Previous Page of LocationAreas",
			CallBack:callback.CallBackMapBack,
		},
	}
}

