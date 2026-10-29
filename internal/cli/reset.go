package cli

import (
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.DB.Reset(s.CTX)
	if err != nil {
		return fmt.Errorf("reset failed: %w", err)
	} else {
		fmt.Println("Database reset!")
	}
	return err
}
