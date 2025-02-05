package main

import (
	"fmt"
	"os"
)

func main() {
	StartRepl(os.Stdin, os.Stdout)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
