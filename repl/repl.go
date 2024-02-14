package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/GiantSnowMonkey/pokedexcli/commands"
)

func StartRepl() {
	for {
		fmt.Print("Pokedexcli > ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		commandsMap := commands.GetCommands()
		command, ok := commandsMap[input.Text()]
		if !ok {
			fmt.Println("command not found!")
			continue
		}
		err := command.Callback()
		if err != nil {
			log.Fatal(err)
		}
	}
}
