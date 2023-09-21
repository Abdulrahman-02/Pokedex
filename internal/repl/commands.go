package repl

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("===== Pokedex HELP =====")
	fmt.Println("Commands:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s\t\t%s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandClear() error {
	fmt.Print("\033[H\033[2J")
	return nil
}
