package main

import (
	"fmt"
	"os"
)

func main() {
	StartRepl(os.Stdin, os.Stdout)
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
