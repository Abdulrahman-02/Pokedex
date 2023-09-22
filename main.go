package main

import (
	"time"

	"github.com/Abdulrahman-02/Pokedex/internal/api"
	"github.com/Abdulrahman-02/Pokedex/internal/repl"
)

func main() {
	Client := api.NewClient(10 * time.Second)
	c := &repl.Config{
		ApiClient: Client,
	}
	repl.StartRepl(c)
}
