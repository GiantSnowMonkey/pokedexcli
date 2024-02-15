package main

import (
	"time"

	"github.com/GiantSnowMonkey/pokedexcli/internal/pokeapi"
	"github.com/GiantSnowMonkey/pokedexcli/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &repl.Config{
		PokeapiClient: pokeClient,
	}

	repl.StartRepl(cfg)
}
