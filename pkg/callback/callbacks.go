package callback

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"

	customTypes "github.com/humbl3LilaS/pokedexcli/pkg/types"
)


func CallbackExit(arg customTypes.CmdArg) error {
	fmt.Println("Exiting Pokedex CLI....")
	os.Exit(0)
	return nil
}


func CallbackMap(arg customTypes.CmdArg) error {

	resp, err := arg.AppConf.ApiClient.ListLocationAreas(arg.AppConf.NextLocUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Location Areas: ")

	for _, loc := range resp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	arg.AppConf.NextLocUrl = resp.Next
	arg.AppConf.PrevLocUrl = resp.Previous

	return nil
}

func CallbackMapPrevious(arg customTypes.CmdArg) error {

	if arg.AppConf.PrevLocUrl == nil {
		return errors.New("You are on the first page")
	}

	resp, err := arg.AppConf.ApiClient.ListLocationAreas(arg.AppConf.PrevLocUrl)

	if err != nil {
		log.Fatal(err)
		return err
	}
	
	resp.PrintDetail()

	arg.AppConf.NextLocUrl = resp.Next
	arg.AppConf.PrevLocUrl = resp.Previous

	return nil
}

func CallbackExplore(arg customTypes.CmdArg) error {

	args := arg.Args

	resp , err := arg.AppConf.ApiClient.GetLocationArea(args[0])

	if err != nil {
		log.Fatal(err)
		return err
	}

	resp.PrintDetail()


	return nil
}

func CallbackCatch(arg customTypes.CmdArg) error {

	args := arg.Args

	resp, err := arg.AppConf.ApiClient.CacthPokemon(args[0])

	if err != nil {
		log.Fatal(err)
		return err
	}

	const treshold int = 50

	badFaith := rand.Intn(resp.BaseExperience)

	if badFaith > treshold {
		return fmt.Errorf("Failed To Catch Pokemon. %s escaped.", resp.Name)
	}


	fmt.Printf("Successfully catched a %s.\n", resp.Name)

	return nil
}
