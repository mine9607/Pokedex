package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
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
}

func helpMessage(config *config) error {
	// Note: iterate over the commands map to generate the help message that automatically updates
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage: \n\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// NOTE: need to add back in the update urls function call
func getLocations(c *config) error {
	data, err := c.pokeapiClient.GetLocations(c.next_URL)
	if err != nil {
		return err
	}
	for _, area := range data.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func printAreas() {

}

func mapBack(config *config) error {
	if config.previous_URL == "" {
		fmt.Printf("You're on the first page\n")
	} else {
		config.pokeapiClient.GetLocations(config.previous_URL)
	}
	return nil
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
