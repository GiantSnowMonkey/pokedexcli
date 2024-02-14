package commands

import "os"

func commandExit() error {
	os.Exit(0)
	return nil
}
