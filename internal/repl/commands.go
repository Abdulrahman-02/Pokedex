package repl

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/Abdulrahman-02/Pokedex/internal/api"
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
	location, err := c.ApiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s ... \n", name)
	fmt.Println("Found Pokemon: ")

	for _, location := range location.PokemonEncounters {
		fmt.Println(location.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemon name is mandatory")
	}
	name := args[0]
	pokemon, err := c.ApiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s ... \n", name)

	if chance > 50 {
		fmt.Printf("%s ran away! \n", pokemon.Name)
		return nil
	}
	api.PokemonCaught[pokemon.Name] = pokemon
	fmt.Printf("%s was caught! \n", pokemon.Name)

	return nil
}

func commandInspect(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("pokemon name is mandatory")
	}
	name := args[0]
	details, ok := api.PokemonCaught[name]
	if !ok {
		fmt.Printf("You don't have %s in your Pokedex \n", name)
		return nil
	}
	fmt.Printf("Name: %s \n", details.Name)
	fmt.Printf("Height: %d \n", details.Height)
	fmt.Printf("Weight: %d \n", details.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range details.Stats {
		fmt.Printf("	-%s: %d \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types")
	for _, typeInfo := range details.Types {
		fmt.Println("	-", typeInfo.Type.Name)
	}
	return nil
}

func commandPokedex(c *Config, args ...string) error {
	pokemons := api.PokemonCaught
	fmt.Println("Your Pokedex")
	for _, pokemon := range pokemons {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
