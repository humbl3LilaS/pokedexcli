package pokeapi

import (
	"net/http"
	"time"
)

// TODO: Relace with ENV variable later
const baseURL = "http://localhost/api/v2"

func NewClient() Client {
	return Client{
		HttpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
