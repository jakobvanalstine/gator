package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

// Struct that holds fields of an RSS feed.
// NOTICE: This code is boilerplate provided by
// https://www.boot.dev/courses/build-blog-aggregator-golang
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

// Struct that holds fields of an RSS item.
// NOTICE: This code is boilerplate provided by
// https://www.boot.dev/courses/build-blog-aggregator-golang
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

var client = &http.Client{Timeout: 5 * time.Second}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &feed, err
	}
	req.Header.Set("User-Agent", "gator")
	req.Header.Set("Accept", "application/xhtml+xml,application/xml")

	resp, err := client.Do(req)
	if err != nil {
		return &feed, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &feed, err
	}

	if err := xml.Unmarshal(dat, &feed); err != nil {
		return &feed, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, err
}
