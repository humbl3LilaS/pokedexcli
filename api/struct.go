package pokeapi

import (
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
