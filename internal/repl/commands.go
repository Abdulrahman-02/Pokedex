package repl

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(*Config, ...string) error {
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

func commandExit(*Config, ...string) error {
	os.Exit(0)
	return nil
}

func commandClear(*Config, ...string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

func commandMap(c *Config, args ...string) error {
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

func commandMapb(c *Config, args ...string) error {
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

func commandExplore(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("location name is mandatory")
	}
	name := args[0]
	pokemons, err := c.ApiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s ... \n", name)
	fmt.Println("Found Pokemon: ")

	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
