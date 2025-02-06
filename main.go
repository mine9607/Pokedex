package main

import (
	"os"
	"time"

	"github.com/mine9607/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	config := &config{
		pokeapiClient: pokeClient,
		next_URL:      "",
		curr_URL:      "",
		previous_URL:  "",
	}

	StartRepl(os.Stdin, os.Stdout, config)
}
