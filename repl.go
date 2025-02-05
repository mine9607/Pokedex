package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
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
		callback:    displayMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Returns the previous 20 locations in the Pokemon world",
		callback:    mapBack,
	}
}

func StartRepl(in io.Reader, out io.Writer) {
	// create a new scanner to wait for input from the CLI
	scanner := bufio.NewScanner(in)

	// Create a config object
	c := NewConfig()
	// create an infinite loop to wait for input execution (text + ENTER)

	for {
		fmt.Print("Pokedex > ")

		// scan the input text
		if !scanner.Scan() {
			break
		}

		// get the input text
		input := scanner.Text()

		// clean the input and convert to slices of lowercase words
		words, err := cleanInput(input)
		if err != nil {
			fmt.Fprintln(out, "Error", err)
			continue
		}

		command := words[0]

		if cmd, ok := commands[command]; ok {
			// call the callback function and print any errors that are returned
			err := cmd.callback(c)
			if err != nil {
				fmt.Fprintf(out, "Error executing command %s: %v\n:", command, err)
			}
		} else {
			fmt.Fprintf(out, "Unknown command: %s\n", command)

		}

		io.WriteString(out, "\n")
	}
}

func cleanInput(text string) ([]string, error) {

	words := strings.Fields(strings.ToLower(text))
	if len(words) == 0 {
		return nil, fmt.Errorf("Input string must be valid text with length > 0")
	}
	return words, nil
}

func helpMessage(config *Config) error {
	// Note: iterate over the commands map to generate the help message that automatically updates
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage: \n\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func displayMap(c *Config) error {
	// using the command map to call a fetch to location-area

	// should display the names of 20 location areas in the Pokemon World
	// make a request to the PokeAPI location-area endpoint
	// each call should display the next_20 locations (using the config.Next_URL I imagine)

	base_url := "https://pokeapi.co/api/v2/location-area"

	// update the Next_URL in the config
	//	Next_URL(url, c)

	// function call which passes the needed url to the GET function which makes the request and returns the data
	// Check if a next_url exists: if not use base_url else use next_url
	if c.Next_URL == "" {
		getAreas(base_url, c)
	} else {
		getAreas(c.Next_URL, c)
	}

	return nil
}

func getAreas(url string, c *Config) error {
	// Makes a get request to location-area endpoint

	// Make a get request to the base_url or c.Next_URL
	data, err := GET(url)
	if err != nil {
		return fmt.Errorf("Error fetching location data:", err)
	}
	// Update the config instance URLs
	//fmt.Printf("Prev URL: %s\nNext URL: %s\n", c.Previous_URL, c.Next_URL)
	if data.Previous == "" {
		UpdateConfigURLs("", data.Next, c)
	} else {
		UpdateConfigURLs(data.Previous, data.Next, c)
	}
	//fmt.Printf("Prev URL: %s\nNext URL: %s\n", c.Previous_URL, c.Next_URL)

	//.areas := data["results"].([]map[string]interface{})
	fmt.Println(data.Next)
	printAreas(data.Results)
	return nil
}

func printAreas(areas []LocationAreasType) {
	for _, area := range areas {
		fmt.Println(area.Name)
	}
}

func mapBack(c *Config) error {
	if c.Previous_URL == "" {
		fmt.Printf("You're on the first page\n")
	} else {
		getAreas(c.Previous_URL, c)
	}
	return nil
}
