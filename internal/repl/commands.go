package repl

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(*Config) error {
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

func commandExit(*Config) error {
	os.Exit(0)
	return nil
}

func commandClear(*Config) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

func commandMap(c *Config) error {
	locations, err := c.ApiClient.GetLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = locations.Next
	c.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(c *Config) error {
	if c.prevLocationsURL == nil {
		return errors.New("no previous locations")
	}
	locations, err := c.ApiClient.GetLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = locations.Next
	c.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
