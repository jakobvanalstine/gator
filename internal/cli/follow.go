package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected one argument: a URL")
	}
	url := cmd.args[0]

	user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUserName)
	if err != nil {
		return errors.New("please log in first")
	}

	feed, err := s.DB.GetFeedForURL(context.Background(), url)
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

	_, err = s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(user.Name, "now follows", feed.Name)
	return nil
}
