package cli

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.DB.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("reset failed: %w", err)
	} else {
		fmt.Println("Database reset!")
	}
	return err
}
