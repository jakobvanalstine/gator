package cli

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if cmd.args == nil || len(cmd.args) != 1 {
		return errors.New("expected a duration string for time between requests")
	}
	time_between_reqs := cmd.args[0]

	d, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	fmt.Println("Collecting feeds every", time_between_reqs)

	if err := scrapeFeeds(s); err != nil {
		return err
	}
	for {
		fmt.Printf("\nSleeping...\n\n")
		select {
		case <-s.CTX.Done():
			return s.CTX.Err()
		case <-ticker.C:
			if err := scrapeFeeds(s); err != nil {
				return err
			}
		}
	}

	return nil
}
