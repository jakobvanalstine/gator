package cli

import (
	"context"
	"fmt"

	"github.com/jakobvanalstine/gator/internal/rss"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	f, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", f)
	return err
}
