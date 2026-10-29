package cli

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected one argument: a URL")
	}
	url := cmd.args[0]

	feed, err := s.DB.GetFeedForURL(s.CTX, url)
	if err != nil {
		return errors.New("please add the feed first")
	}

	createdAt := time.Now()
	updatedAt := createdAt
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(s.CTX, params)
	if err != nil {
		return err
	}

	fmt.Println(user.Name, "now follows", feed.Name)
	return nil
}
