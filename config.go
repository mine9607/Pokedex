package main

import "github.com/mine9607/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	next_URL      string
	previous_URL  string
}

func UpdateConfigURLs(prev_url string, next_url string, c *config) {
	// update the previousURL to Next_URL and update Next_URL to input url then return that url
	c.previous_URL = prev_url
	c.next_URL = next_url
}
