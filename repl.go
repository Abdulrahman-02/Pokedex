package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl implements a simple command-line interface (CLI) for a Pokedex application.
func startRepl() {
	fmt.Println("===== Pokedex Unix CLI =====")

	for {
		fmt.Print("pokedex > ")

		// Read the command line
		reader := bufio.NewReader(os.Stdin)

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
			err := command.function()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Command not found")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	function    func() error
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
	}
}
