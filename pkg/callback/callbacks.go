package callback

import (
	"errors"
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

	resp, err := conf.ApiClient.ListLocationAreas(conf.NextLocUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Location Areas: ")

	for _, loc := range resp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	conf.NextLocUrl = resp.Next
	conf.PrevLocUrl = resp.Previous

	return nil
}

func CallBackMapBack(conf *customTypes.AppConfig) error {

	if conf.PrevLocUrl == nil {
		return errors.New("You are on the first page")
	}

	resp, err := conf.ApiClient.ListLocationAreas(conf.PrevLocUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Location Areas: ")

	for _, loc := range resp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	conf.NextLocUrl = resp.Next
	conf.PrevLocUrl = resp.Previous

	return nil
}
