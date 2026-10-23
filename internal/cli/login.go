package cli

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected a username")
	}
	username := cmd.args[0]

	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return errors.New("user does not exist")
	}

	if err := s.CFG.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("logged in as", user.Name)

	return nil
}
