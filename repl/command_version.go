package repl

import "fmt"

func commandVersion(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Pokedexcli v0.0.1")
	return nil
}
