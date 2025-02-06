package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ExploreAreaResponse struct {
}

func (c *Client) GetLocations(to_URL string) (LocationsResponse, error) {
	url := ""
	base_url := "https://pokeapi.co/api/v2/location-area"

	//
	if to_URL != "" {
		url = to_URL
	} else {
		url = base_url
	}

	// check cache for the url before making a fetch
	areas := LocationsResponse{}
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &areas)
		if err != nil {
			return areas, err
		}

		return areas, nil
	}

	// make a GET request to url and return the
	areas, err := c.GET(url)
	if err != nil {
		return areas, err // areas will be LocationsResponse{} if err != nil
	}

	return areas, nil

}

func (c *Client) GetAreaData(URL string) (ExploreAreaResponse, error) {
	// similar to GetLocations but the url will add an {id or name} to the url
	// this will be the second parameter passed to the CLI

	return ExploreAreaResponse{}, nil
}

// The GET method assigns a get request to the client so the client can make a get requet giving access to the cache
func (c *Client) GET(url string) (LocationsResponse, error) {

	response := LocationsResponse{}
	res, err := http.Get(url)
	if err != nil {
		return response, fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return response, fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, fmt.Errorf("error unMarshaling response body: %w", err)
	}

	c.cache.Add(url, data)

	return response, nil
}
