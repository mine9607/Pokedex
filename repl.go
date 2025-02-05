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
	callback    func() error
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
}

func StartRepl(in io.Reader, out io.Writer) {
	// create a new scanner to wait for input from the CLI
	scanner := bufio.NewScanner(in)

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
			err := cmd.callback()
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

func helpMessage() error {
	// Note: iterate over the commands map to generate the help message that automatically updates
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage: \n\n")

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
