package cli

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var userForID = map[uuid.UUID]string{}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	n := len(feeds)
	for i, f := range feeds {
		fmt.Println("Name:", f.Name)
		fmt.Println("URL:", f.Url)

		username, ok := userForID[f.ID]
		if !ok {
			user, err := s.DB.GetUserForID(context.Background(), f.UserID)
			if err != nil {
				return err
			}
			userForID[f.ID] = user.Name
			username = user.Name
		}

		fmt.Println("Created by:", username)

		if i < n-1 {
			fmt.Println()
		}
	}

	return nil
}
