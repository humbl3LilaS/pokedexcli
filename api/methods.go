package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResp{}, nil
	}

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		return LocationAreaResp{}, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("Bad Status Code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreas := LocationAreaResp{}

	err = json.Unmarshal(data, &locationAreas)

	if err != nil {
		return LocationAreaResp{}, err
	}

	return  locationAreas, nil
}
