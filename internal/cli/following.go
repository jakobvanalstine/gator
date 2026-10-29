package cli

import (
	"fmt"

	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ffs, err := s.DB.GetFeedFollowsForUser(s.CTX, user.Name)
	if err != nil {
		return err
	}

	for _, ff := range ffs {
		fmt.Println("*", ff.FeedName)
	}

	return nil
}
