package pokeapi

import (
	"fmt"
	"net/http"

	pokecache "github.com/humbl3LilaS/pokedexcli/cache"
)


type Client struct {
	Cache pokecache.Cache
	HttpClient http.Client
}

type LocationAreaResp struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}


func (loc LocationAreaResp) PrintDetail() {
	fmt.Println("Location Areas: ")

	for _, loc := range loc.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

}


func (loc LocationArea) PrintDetail() {
	fmt.Println("Location Infos: ")

	fmt.Printf(" - name: %s\n", loc.Location.Name)

	fmt.Println("Pokemon Encounters: ")

	for _, val := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", val.Pokemon.Name)
	}

	fmt.Println("Encounter Methods: ")

	for _,val := range loc.EncounterMethodRates {
		fmt.Printf(" - %s\n", val.EncounterMethod.Name)
		for _, ver := range val.VersionDetails{
			fmt.Printf("  - Version: %s, Rate: %d\n", ver.Version.Name, ver.Rate)
		}
	}

}
