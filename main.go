package main

//go:generate sqlboiler --wipe psql -o adapters/repository/models

import (
	"cryptotrade/cmd"
)

func main() {
	cmd.Execute()
}
