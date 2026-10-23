package cli

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.CFG.CurrentUserName {
			fmt.Println("*", u.Name, "(current)")
		} else {
			fmt.Println("*", u.Name)
		}
	}

	return nil
}
