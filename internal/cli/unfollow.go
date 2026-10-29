package cli

import (
	"errors"

	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected the name of a feed")
	}
	url := cmd.args[0]

	feed, err := s.DB.GetFeedForURL(s.CTX, url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.DB.DeleteFeedFollow(s.CTX, params); err != nil {
		return err
	}

	return nil
}
