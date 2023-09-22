package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Abdulrahman-02/Pokedex/internal/api"
)

// startRepl implements a simple command-line interface (CLI) for a Pokedex application.
func StartRepl(c *Config) {
	fmt.Println("===== Pokedex Unix CLI =====")

	// Read the command line
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")

		// Read the input until the first newline
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Trim whitespace from the command string
		cmdString = strings.TrimSpace(cmdString)

		// Handle the execution of the command
		command, exists := getCommands()[cmdString]

		if exists {
			err := command.function(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Command not found")
		}
	}
}

type Config struct {
	ApiClient        api.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	function    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the application",
			function:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			function:    commandClear,
		},
		"help": {
			name:        "help",
			description: "Show help",
			function:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			function:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations",
			function:    commandMapb,
		},
	}
}
