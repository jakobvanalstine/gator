# Gator

## Summary

Gator is an RSS feed aggregator. I wrote it while taking Boot.dev's course [Build a Blog Aggregator in Go](https://www.boot.dev/courses/build-blog-aggregator-golang). 

## Setup

Install the latest Go toolchain.

Create a PostgreSQL database. Set it up using the files in `sql/schema/`. If you don't know how, I'll point you to [Goose](https://github.com/pressly/goose), the database migration tool I use during development.

Your home directory will need a `.gatorconfig.json` file with a PostgreSQL URI:

```json
{
  "db_url": "postgres://<user>:<password>@<hostspec>?sslmode=disable"
}
```

For example, if your PostgreSQL username is `user`, your PostgreSQL password is `password`, and your database is hosted locally on port 5432, the contents would be:


```json
{
  "db_url": "postgres://user:password@localhost:5432?sslmode=disable"
}
```

You can install Gator itself using

```sh
go install ...
```

Having done so, first register a user:

```sh
gator register <username>
```

This will automatically log you into Gator as that user. You can switch between users with the command

```sh
gator login <user>
```

The next command lists all registered users and also specifies who the current user is:

```sh
gator users
```

## Usage

Add a feed you'd like to follow:

```sh
gator addfeed <name> <url>
```

You can add more than one feed, but only one at a time. To see all feeds, use the command

```sh
gator feeds
```

The current user will automatically begin following a feed you add. Another user (once they are the current user) can manually follow that feed with the command

```sh
gator follow <url>
```

This lets you unfollow feeds:

```sh
gator unfollow <url>
```

To see which feeds the current user is following, use this command:

```sh
gator following
```

The following command tells Gator to aggregate posts from the feeds followed by the current user:

```sh
gator agg <duration string>
```

A duration string is any string parsable by Go's `std/time.ParseDuration` function (e.g., `30s`, `20m`, `1h`). You should use `Ctrl+c`/`Cmd+c` to stop aggregating.

When you want to browse aggregated posts, input

```sh
gator browse [limit]
```

A limit is an optional number representing the maximum number of posts to retrieve. It defaults to `2`.

This command will print the title and URL of each aggregated post. To read the post, you'll have to follow the link.
