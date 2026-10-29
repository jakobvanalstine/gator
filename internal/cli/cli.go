package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jakobvanalstine/gator/internal/config"
	"github.com/jakobvanalstine/gator/internal/database"
	"github.com/jakobvanalstine/gator/internal/rss"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
)

func NewState() *state {
	return &state{}
}

func GetCommands() *commands {
	c := &commands{
		byName: map[string]func(*state, command) error{},
	}
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("agg", handlerAgg)
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("users", handlerUsers)
	return c
}

func MakeCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}

func (c *commands) Run(s *state, cmd command) error {
	handler, ok := c.byName[cmd.name]
	if !ok {
		return fmt.Errorf("no command registered as: %v\n", cmd.name)
	}

	return handler(s, cmd)
}

type state struct {
	CFG *config.Config
	DB  *database.Queries
	CTX context.Context
}

type command struct {
	name string
	args []string
}

type commands struct {
	byName map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.byName[name] = f
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.DB.GetUser(s.CTX, s.CFG.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}

func scrapeFeeds(s *state) error {
	dbFeed, err := s.DB.GetNextFeedToFetch(s.CTX)
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		ID: dbFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	if err := s.DB.MarkFeedFetched(s.CTX, params); err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(s.CTX, dbFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		now := time.Now().UTC()
		params := database.AddPostParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Title:     item.Title,
			Url:       item.Link,
			FeedID:    dbFeed.ID,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
		}
		time, err := tryTimeParse(item.PubDate)
		params.PublishedAt = sql.NullTime{
			Time:  time,
			Valid: err == nil, // Note the operator used.
		}

		_, err = s.DB.AddPost(s.CTX, params)
		if err != nil {
			pqErr := new(pq.Error)
			if errors.As(err, &pqErr) {
				switch pqErr.Code {
				case pqerror.UniqueViolation:
					// Scraping feeds is bound to turn up posts we already have.
					// Ignore them instead of crashing the scraper.
				default:
					fmt.Fprintln(os.Stderr, "WARNING:", err.Error())
				}
			} else {
				fmt.Fprintln(os.Stderr, "WARNING:", err.Error())
			}
		}
	}

	return nil
}

var (
	tryTimeParseCachedIndex = -1
	tryTimeParseLayouts     = []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
	}
)

func tryTimeParse(value string) (time.Time, error) {
	if tryTimeParseCachedIndex != -1 && tryTimeParseCachedIndex < len(tryTimeParseLayouts) {
		layout := tryTimeParseLayouts[tryTimeParseCachedIndex]
		time, err := time.Parse(layout, value)
		if err != nil {
			// Do nothing.
		} else {
			return time, nil
		}
	}

	for i, layout := range tryTimeParseLayouts {
		time, err := time.Parse(layout, value)
		if err != nil {
			continue
		}

		// A layout worked.
		tryTimeParseCachedIndex = i
		return time, nil
	}

	// No layout worked.
	return time.Time{}, errors.New("failed to parse time")
}
