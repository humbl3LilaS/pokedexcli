package callback

import (
	"fmt"
	"log"
	"os"
	customTypes "github.com/humbl3LilaS/pokedexcli/pkg/types"
)


func CallbackExit(conf *customTypes.AppConfig) error {
	fmt.Println("Exiting Pokedex CLI....")
	os.Exit(0)
	return nil
}


func CallbackMap(conf *customTypes.AppConfig) error {

	resp, err := conf.ApiClient.ListLocationAreas()

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
