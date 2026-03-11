package main

import (
	pokeapi "github.com/humbl3LilaS/pokedexcli/api"
	"github.com/humbl3LilaS/pokedexcli/cmd"
	customTypes "github.com/humbl3LilaS/pokedexcli/pkg/types"
)



func main() {
	appConf := customTypes.AppConfig{
		ApiClient: pokeapi.NewClient(),
	}

	cmd.StartRepl(&appConf)
}
