# Pokedex

A simple command-line interface (CLI) for a Pokedex application written in Go.

## Usage

This CLI provides various commands to interact with the Pokedex. Here are the available commands:

- `exit`: Exit the application.
- `clear`: Clear the screen.
- `help`: Show help.
- `map`: Display the names of 20 location areas in the Pokemon world.
- `mapb`: Display the previous 20 locations.
- `explore`: Display Pokemon names in an area.
- `catch`: Catch Pokemon.
- `inspect`: Inspect Pokemon.
- `pokedex`: Show Pokemon in the Pokedex.

To run the CLI, simply execute the following command:

```bash
go run main.go
```

## TO-DO

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

Feel free to contribute to this project or make any improvements you find valuable. Happy coding!
