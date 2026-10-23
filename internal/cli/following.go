package cli

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	ffs, err := s.DB.GetFeedFollowsForUser(context.Background(), s.CFG.CurrentUserName)
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
