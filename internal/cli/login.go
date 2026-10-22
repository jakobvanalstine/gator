package cli

import (
	"context"
	"errors"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected a single argument: a username")
	}
	username := cmd.args[0]

	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return errors.New("user does not exist")
	}

	if err := s.CFG.SetUser(user.Name); err != nil {
		return err
	}
	log.Println("username set to", user.Name)

	return nil
}
