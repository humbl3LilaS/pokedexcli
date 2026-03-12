package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/humbl3LilaS/pokedexcli/cache"
)

// TODO: Relace with ENV variable later
const baseURL = "http://localhost/api/v2"

func NewClient(interval time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
		Cache: pokecache.NewCache(),
	}
}
