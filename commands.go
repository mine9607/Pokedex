package main

import (
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, params []string) error
}

var commands = make(map[string]cliCommand)

func init() {
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    helpMessage,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Display next 20 locations in the Pokemon world",
		callback:    getLocations,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Returns the previous 20 locations in the Pokemon world",
		callback:    mapBack,
	}
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Checks for Pokemon in an area",
		callback:    exploreArea,
	}
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Attempt to catch a Pokemon in the area",
		callback:    catchPokemon,
	}
	commands["cache"] = cliCommand{
		name:        "cache",
		description: "Returns current cache",
		callback:    checkCache,
	}
}

func helpMessage(config *config, params []string) error {
	// Note: iterate over the commands map to generate the help message that automatically updates
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage: \n\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// NOTE: need to add back in the update urls function call
func getLocations(config *config, params []string) error {
	data, err := config.pokeapiClient.GetLocations(config.next_URL)
	if err != nil {
		return err
	}

	UpdateConfigURLs(data.Previous, data.Next, config, "f")

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	fmt.Printf("PREV URL: %s\n", config.previous_URL)
	fmt.Printf("CURR URL: %s\n", config.curr_URL)
	fmt.Printf("NEXT URL: %s\n", config.next_URL)

	return nil
}

func printAreas() {

}

// NOTE: mapBack not showing FOUND IN CACHE AFTER INITIAL NAVIGATION BACK to first page
// EXAMPLE: calling "map" "map" "mapb" doesn't show that the second page was found in Cache
// NOTE: This is because the returned url contains the query params for setting the returned 20 locations
// base_url: https://pokeapi.co/api/v2/location-area
// returned base_url: https://pokeapi.co/api/v2/location-area?offset=0&limit=20
func mapBack(config *config, params []string) error {
	if config.previous_URL == "" {
		fmt.Printf("You're on the first page\n")
		return nil
	}

	data, err := config.pokeapiClient.GetLocations(config.previous_URL)
	if err != nil {
		return err
	}

	UpdateConfigURLs(data.Previous, data.Next, config, "b")

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	fmt.Printf("PREV URL: %s\n", config.previous_URL)
	fmt.Printf("CURR URL: %s\n", config.curr_URL)
	fmt.Printf("NEXT URL: %s\n", config.next_URL)

	return nil
}

func commandExit(config *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func exploreArea(config *config, params []string) error {
	// check length of params to ensure area provided
	if len(params) == 0 {
		return fmt.Errorf("Error: Must provide an area to explore")
	}

	// NOTE: the split of the parameters is not allowing "-" in names
	area := params[0]

	// create the url to explore and pass to the clientMethd ExploreArea
	base_url := "https://pokeapi.co/api/v2/location-area"
	explore_url := base_url + "/" + area

	data, err := config.pokeapiClient.ExploreArea(explore_url)
	if err != nil {
		return err
	}

	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil

}

func catchPokemon(config *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Error: Must include the name of a pokemon")
	}

	pokemon := params[0]

	// Check if pokemon already in Pokedex
	if _, ok := config.pokedex[pokemon]; ok {
		fmt.Printf("%s already in Pokedex!\n Use command 'inspect' to view it.", pokemon)
		return nil
	}

	// create the endpoint
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	// print the following:
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	// Make a Get Request for the Pokemon's stats
	data, err := config.pokeapiClient.CatchPokemon(url)
	if err != nil {
		return err
	}

	// Use the Pokemon's "base experience" to determine the chance of catching it
	// assuming Mewtwo has highest base experience of 340
	experience := data.BaseExperience
	chance := rand.Intn(experience)

	if chance < 40 {
		fmt.Printf("%s got away!", data.Name)
		return nil
	}

	fmt.Printf("%s was caught!!!", pokemon)
	config.pokedex[pokemon] = data
	// Once the Pokemon is caught, add it to the user's Pokedex, (map[string]Pokemon)

	// Test the "catch" command manually

	return nil
}

func checkCache(config *config, params []string) error {
	// add a check for previous_url = "" && next_url = "" in that case there will be no cache yet

	url := config.previous_URL
	cache := config.pokeapiClient.GetCache()
	if data, ok := cache.Get(url); ok {
		fmt.Println(data)
	}
	return nil
}
