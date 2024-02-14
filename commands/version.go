package commands

import "fmt"

func commandVersion() error {
	fmt.Println()
	fmt.Println("Pokedexcli v0.0.1")
	return nil
}
