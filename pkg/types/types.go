package customTypes

import pokeapi "github.com/humbl3LilaS/pokedexcli/api"

type CliCommand struct {
	Name string
	Desctiption string
	CallBack func(CmdArg) error
	MaxArg int
}

type CmdArg struct {
	AppConf *AppConfig
	Args []string
}

type AppConfig struct{
	ApiClient pokeapi.Client
	NextLocUrl *string
	PrevLocUrl *string
}

