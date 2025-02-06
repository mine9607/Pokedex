package main

import (
	"fmt"
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
	/* check cache
	-> if cache data is not null:
		1) check for area name in string(data.Results)



	*/

	// note: if we haven't called map then previous_URL = "" and next_URL = ""

	// in GetAreaData we need to validate that previous_URL exists (so we can build the correct URL)

	// in GetAreaData we need to validate the url path exists in cache and that there is
	//url := config.previous_URL + params[0]
	//data, err := config.pokeapiClient.GetAreaData(url)

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
