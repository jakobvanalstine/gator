package cli

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return errors.New("expected arguments")
	}

	if len(cmd.Args) != 1 {
		return errors.New("expected a single argument: a username")
	}
	username := cmd.Args[0]

	if err := s.Config.SetUser(username); err != nil {
		return err
	}
	fmt.Println("username is set to:", username)

	return nil
}
