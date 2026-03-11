package customTypes

import pokeapi "github.com/humbl3LilaS/pokedexcli/api"

type CliCommand struct {
	Name string
	Desctiption string
	CallBack func(*AppConfig) error
}

type AppConfig struct{
	ApiClient pokeapi.Client
	NextLocUrl *string
	PrevLocUrl *string
}

