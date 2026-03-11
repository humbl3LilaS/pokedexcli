package callback

type CliCommand struct {
	Name string
	Desctiption string
	CallBack func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name: "help",
			Desctiption: "Print the help menu of Pokedex CLI.",
			CallBack: callbackHelp,
		},
		"exit": {
			Name: "exit",
			Desctiption: "Turnoff the Pokedex CLI.",
			CallBack: callbackExit,
		},
		"map": {
			Name: "map",
			Desctiption: "Output the LocationAreas",
			CallBack: callbackMap,
		},
	}
}

