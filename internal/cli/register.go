package cli

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected a username")
	}
	username := cmd.args[0]

	now := time.Now().UTC()
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	}

	user, err := s.DB.CreateUser(s.CTX, params)
	if err != nil {
		return err
	}

	if err := s.CFG.SetUser(username); err != nil {
		return err
	}

	fmt.Println("created user", user.Name)

	return err
}
