package cli

import (
	"fmt"

	"github.com/jakobvanalstine/gator/internal/config"
	"github.com/jakobvanalstine/gator/internal/database"
)

func NewState() *state {
	return &state{}
}

func GetCommands() *commands {
	c := &commands{
		byName: map[string]func(*state, command) error{},
	}
	c.register("addfeed", handlerAddFeed)
	c.register("agg", handlerAgg)
	c.register("feeds", handlerFeeds)
	c.register("follow", handlerFollow)
	c.register("following", handlerFollowing)
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	return c
}

func MakeCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}

func (c *commands) Run(s *state, cmd command) error {
	handler, ok := c.byName[cmd.name]
	if !ok {
		return fmt.Errorf("no command registered as: %v\n", cmd.name)
	}

	return handler(s, cmd)
}

type state struct {
	CFG *config.Config
	DB  *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	byName map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.byName[name] = f
}
