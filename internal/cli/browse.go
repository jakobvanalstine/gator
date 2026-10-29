package cli

import (
	"fmt"
	"strconv"

	"github.com/jakobvanalstine/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) > 0 {
		i, err := strconv.ParseInt(cmd.args[0], 0, 32)
		if err != nil {
			return err
		}

		limit = int32(i)
	}

	params := database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: limit,
	}

	posts, err := s.DB.GetPostsForUser(s.CTX, params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("\n%+v\n", post)
	}

	return nil
}
