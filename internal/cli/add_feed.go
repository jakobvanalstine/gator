package cli

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if cmd.args == nil || len(cmd.args) != 2 {
		return errors.New("expected two arguments: a name, a URL")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	createdAt := time.Now()
	updatedAt := createdAt
	params := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.DB.AddFeed(s.CTX, params)
	if err != nil {
		if !isUniqueViolation(err) {
			return err
		}
	}
	fmt.Println("created feed", feed.Name)

	newCMD := command{
		name: "follow",
		args: []string{url},
	}
	if err := handlerFollow(s, newCMD, user); err != nil {
		if !isUniqueViolation(err) {
			return err
		}
	}

	return nil
}
