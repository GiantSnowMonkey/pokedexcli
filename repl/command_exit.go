package repl

import "os"

func commandExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}
