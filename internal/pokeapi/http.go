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
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// The GET method assigns a get request to the client so the client can make a get requet giving access to the cache
// func (c *Client) GET(url string) (LocationsResponse, error) {
func (c *Client) GET(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		//return response, fmt.Errorf("error making request: %w", err)
		return nil, fmt.Errorf("Error making request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		//return response, fmt.Errorf("error reading response body: %w", err)
		return nil, fmt.Errorf("Error reading the response body: %w", err)
	}

	c.cache.Add(url, data)
	return data, nil
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
		fmt.Println("FOUND LOCATION IN CACHE!!!!!!!")
		err := json.Unmarshal(val, &areas)
		if err != nil {
			return areas, err
		}

		return areas, nil
	}

	// make a GET request to url and return the
	data, err := c.GET(url)
	if err != nil {
		return areas, err // areas will be LocationsResponse{} if err != nil
	}
	if err := json.Unmarshal(data, &areas); err != nil {
		return areas, err
	}

	return areas, nil

}

func (c *Client) ExploreArea(area_url string) (ExploreAreaResponse, error) {
	// similar to GetLocations but the url will add an {id or name} to the url

	// Check cache for existing url
	exploreArea := ExploreAreaResponse{}
	if val, ok := c.cache.Get(area_url); ok {
		fmt.Println("FOUND AREA IN CACHE!!!")
		err := json.Unmarshal(val, &exploreArea)
		if err != nil {
			return exploreArea, err
		}

		return exploreArea, nil
	}

	// if not found in cache make a GET request
	data, err := c.GET(area_url)
	if err != nil {
		return exploreArea, err
	}

	if err := json.Unmarshal(data, &exploreArea); err != nil {
		return exploreArea, fmt.Errorf("Error unmarshaling area data: %w", err)
	}

	return exploreArea, nil
}

func (c *Client) CatchPokemon(url string) (Pokemon, error) {

	// check Pokedex
	// note: add a check to see if pokemon in current area

	pokemon := Pokemon{}

	// make Get request
	data, err := c.GET(url)
	if err != nil {
		return pokemon, fmt.Errorf("Error fetching pokemon stats: %w", err)
	}

	// attempt to unmarshal data
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, fmt.Errorf("Error unmarshaling Pokemon stats: %w", err)
	}

	// return fetched data
	return pokemon, nil
}
