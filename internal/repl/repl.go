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
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// Handle the execution of the command
		command, exists := getCommands()[commandName]

		if exists {
			err := command.function(c, args...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		} else {
			fmt.Fprintln(os.Stderr, "Command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type Config struct {
	ApiClient        api.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	function    func(*Config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "displays pokemon names in area",
			function:    commandExplore,
		},
	}
}
