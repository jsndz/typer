package main

import (
	"os"

	"github.com/jsndz/typer/commands"
)

func main() {
	Args := os.Args

	switch Args[1] {
	case "open":
		commands.Open(Args[2])
	case "create":
		commands.Create(Args[2])
	case "delete":
		commands.Delete(Args[2])
	case "edit":

	}
}
