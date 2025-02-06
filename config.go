package main

import "github.com/mine9607/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	next_URL      string
	curr_URL      string
	previous_URL  string
}

func UpdateConfigURLs(prev_url string, next_url string, c *config, flag string) {
	// note the flag is "f" or "b" for forward or backwards navigation based on the call of map or mapb

	if c.previous_URL == "" && c.next_URL == "" {
		c.curr_URL = "https://pokeapi.co/api/v2/location-area"
	} else if flag == "f" {
		c.curr_URL = c.next_URL
	} else {
		c.curr_URL = c.previous_URL
	}
	// Update previous and next URLs
	c.previous_URL = prev_url
	c.next_URL = next_url
}
