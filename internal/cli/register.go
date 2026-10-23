package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected one argument: a username")
	}
	username := cmd.args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err == nil {
		// OBS: We check err == nil to see if a user by this name already exists.
		return errors.New("user already exists")
	}

	createdAt := time.Now()
	updatedAt := createdAt
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      username,
	}

	if err := s.CFG.SetUser(username); err != nil {
		return err
	}

	user, err := s.DB.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Created user %v\n", user.Name)
	fmt.Printf("%+v\n", user)

	return err
}
