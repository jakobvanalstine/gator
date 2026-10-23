package cli

import (
	"context"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.DB.Reset(context.Background())
	if err != nil {
		log.Println("Failed to reset database!")
	} else {
		log.Println("Database reset!")
	}
	return err
}
