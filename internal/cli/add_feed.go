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
		return err
	}
	fmt.Printf("Created feed %v\n", feed.Name)
	fmt.Printf("%+v\n", feed)

	newCMD := command{
		name: "follow",
		args: []string{url},
	}
	if err := handlerFollow(s, newCMD, user); err != nil {
		// TODO(jakob): If this fails, cleanup should remove the new feed.
		return err
	}

	return nil
}
