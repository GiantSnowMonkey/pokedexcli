package commands

import "fmt"

func commandHelp() error {
	commandsMap := GetCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for i := range commandsMap {
		print(commandsMap[i].name, ": ", commandsMap[i].description, "\n")
	}
	return nil
}
