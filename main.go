package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jakobvanalstine/gator/internal/cli"
	"github.com/jakobvanalstine/gator/internal/config"
	"github.com/jakobvanalstine/gator/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected an argument")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	state := cli.NewState()
	state.CFG = cfg
	state.DB = dbQueries

	commands := cli.GetCommands()

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	name := os.Args[1]
	cmd := cli.MakeCommand(name, args)

	if err := commands.Run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
