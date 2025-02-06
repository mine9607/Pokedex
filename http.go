package main

import (
	"encoding/json"
	"fmt"
	"github.com/mine9607/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type LocationAreaGetResponse struct {
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []LocationAreasType `json:"results"`
}

type LocationAreasType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GET(url string) (*LocationAreaGetResponse, error) {
	// check cache

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	//data := make(map[string]any)

	// NOTE: this is now less generic since we need to pass the
	var response LocationAreaGetResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unMarshaling response body: %w", err)
	}

	return &response, nil
}
