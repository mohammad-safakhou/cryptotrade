package main

import (
	"cryptotrade/cmd"
)
//go:generate sqlboiler --wipe --no-tests psql -o models

func main() {
	cmd.Execute()
}
