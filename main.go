package main

import (
	"log"
	"os"

	"github.com/jakobvanalstine/gator/internal/cli"
	"github.com/jakobvanalstine/gator/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected an argument")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	state := cli.MakeState(cfg)
	commands := cli.GetCommands()

	name := os.Args[1]
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	cmd := cli.MakeCommand(name, args)

	if err := commands.Run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
