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

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
			Ability  any  `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems              []any  `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order any `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		Other struct {
			Home struct {
				FrontShiny       any `json:"front_shiny"`
				FrontFemale      any `json:"front_female"`
				FrontDefault     any `json:"front_default"`
				FrontShinyFemale any `json:"front_shiny_female"`
			} `json:"home"`
			Showdown struct {
				BackShiny        any `json:"back_shiny"`
				BackFemale       any `json:"back_female"`
				FrontShiny       any `json:"front_shiny"`
				BackDefault      any `json:"back_default"`
				FrontFemale      any `json:"front_female"`
				FrontDefault     any `json:"front_default"`
				BackShinyFemale  any `json:"back_shiny_female"`
				FrontShinyFemale any `json:"front_shiny_female"`
			} `json:"showdown"`
			DreamWorld struct {
				FrontFemale  any `json:"front_female"`
				FrontDefault any `json:"front_default"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontShiny   any `json:"front_shiny"`
				FrontDefault any `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				Yellow struct {
					BackGray         any `json:"back_gray"`
					FrontGray        any `json:"front_gray"`
					BackDefault      any `json:"back_default"`
					FrontDefault     any `json:"front_default"`
					BackTransparent  any `json:"back_transparent"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"yellow"`
				RedBlue struct {
					BackGray         any `json:"back_gray"`
					FrontGray        any `json:"front_gray"`
					BackDefault      any `json:"back_default"`
					FrontDefault     any `json:"front_default"`
					BackTransparent  any `json:"back_transparent"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"red-blue"`
			} `json:"generation-i"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackShiny        any `json:"back_shiny"`
						BackFemale       any `json:"back_female"`
						FrontShiny       any `json:"front_shiny"`
						BackDefault      any `json:"back_default"`
						FrontFemale      any `json:"front_female"`
						FrontDefault     any `json:"front_default"`
						BackShinyFemale  any `json:"back_shiny_female"`
						FrontShinyFemale any `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationIi struct {
				Gold struct {
					BackShiny        any `json:"back_shiny"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontDefault     any `json:"front_default"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackShiny        any `json:"back_shiny"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontDefault     any `json:"front_default"`
					FrontTransparent any `json:"front_transparent"`
				} `json:"silver"`
				Crystal struct {
					BackShiny             any `json:"back_shiny"`
					FrontShiny            any `json:"front_shiny"`
					BackDefault           any `json:"back_default"`
					FrontDefault          any `json:"front_default"`
					BackTransparent       any `json:"back_transparent"`
					FrontTransparent      any `json:"front_transparent"`
					BackShinyTransparent  any `json:"back_shiny_transparent"`
					FrontShinyTransparent any `json:"front_shiny_transparent"`
				} `json:"crystal"`
			} `json:"generation-ii"`
			GenerationIv struct {
				Platinum struct {
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"platinum"`
				DiamondPearl struct {
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackShiny        any `json:"back_shiny"`
					BackFemale       any `json:"back_female"`
					FrontShiny       any `json:"front_shiny"`
					BackDefault      any `json:"back_default"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					BackShinyFemale  any `json:"back_shiny_female"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
			} `json:"generation-iv"`
			GenerationIx struct {
				ScarletViolet struct {
					FrontFemale  any `json:"front_female"`
					FrontDefault any `json:"front_default"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationVi struct {
				XY struct {
					FrontShiny       any `json:"front_shiny"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"x-y"`
				OmegarubyAlphasapphire struct {
					FrontShiny       any `json:"front_shiny"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
			} `json:"generation-vi"`
			GenerationIii struct {
				Emerald struct {
					FrontShiny   any `json:"front_shiny"`
					FrontDefault any `json:"front_default"`
				} `json:"emerald"`
				RubySapphire struct {
					BackShiny    any `json:"back_shiny"`
					FrontShiny   any `json:"front_shiny"`
					BackDefault  any `json:"back_default"`
					FrontDefault any `json:"front_default"`
				} `json:"ruby-sapphire"`
				FireredLeafgreen struct {
					BackShiny    any `json:"back_shiny"`
					FrontShiny   any `json:"front_shiny"`
					BackDefault  any `json:"back_default"`
					FrontDefault any `json:"front_default"`
				} `json:"firered-leafgreen"`
			} `json:"generation-iii"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  any `json:"front_female"`
					FrontDefault any `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontShiny       any `json:"front_shiny"`
					FrontFemale      any `json:"front_female"`
					FrontDefault     any `json:"front_default"`
					FrontShinyFemale any `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontFemale  any `json:"front_female"`
					FrontDefault any `json:"front_default"`
				} `json:"icons"`
				BrilliantDiamondShiningPearl struct {
					FrontFemale  any `json:"front_female"`
					FrontDefault any `json:"front_default"`
				} `json:"brilliant-diamond-shining-pearl"`
			} `json:"generation-viii"`
		} `json:"versions"`
		BackShiny        any `json:"back_shiny"`
		BackFemale       any `json:"back_female"`
		FrontShiny       any `json:"front_shiny"`
		BackDefault      any `json:"back_default"`
		FrontFemale      any `json:"front_female"`
		FrontDefault     any `json:"front_default"`
		BackShinyFemale  any `json:"back_shiny_female"`
		FrontShinyFemale any `json:"front_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest any `json:"latest"`
		Legacy any `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
	} `json:"past_stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []any `json:"past_types"`
}
