package main

import (
	command "github.com/carldanley/frigate-clip-cleaner/internal/commands"
)

func main() {
	fcc := command.New()
	fcc.Run()
}
