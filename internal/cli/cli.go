package cli

import "github.com/jakobvanalstine/gator/internal/config"

type state struct {
	Config *config.Config
}

type command struct {
	Name string
	Args []string
}
