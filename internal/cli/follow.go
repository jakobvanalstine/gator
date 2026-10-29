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
		return errors.New("expected one URL")
	}
	url := cmd.args[0]

	feed, err := s.DB.GetFeedForURL(s.CTX, url)
	if err != nil {
		return errors.New("please add the feed and give it a name first")
	}

	now := time.Now()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(s.CTX, params)
	if err != nil {
		if !isUniqueViolation(err) {
			return err
		}
	}

	fmt.Println(user.Name, "now follows", feed.Name)
	return nil
}
