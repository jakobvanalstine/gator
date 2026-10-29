package cli

import (
	"context"
	"fmt"

	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ffs, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	n := len(ffs)
	for i, ff := range ffs {
		fmt.Println("*", ff.FeedName)
		if i < n-1 {
			fmt.Println()
		}
	}

	return nil
}
