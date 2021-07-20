package twitchtracker

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/gocolly/colly/v2"
)

const (
	twitchTrackerBaseUrl = "https://twitchtracker.com/channels/ranking/%s?page=%d"
)

// FindTopStreamers finds top streamers by scraping rankings on https://twitchtracker.com.
func FindTopStreamers(language string, pages int) ([]string, error) {
	c := colly.NewCollector()
	c.AllowURLRevisit = true
	var streamers []string
	c.OnHTML("#channels tr td:nth-child(3) a", func(e *colly.HTMLElement) {
		streamers = append(streamers, e.Text)
	})

	for page := 1; page <= pages; page++ {
		err := backoff.Retry(func() error {
			return c.Visit(fmt.Sprintf(twitchTrackerBaseUrl, language, page))
		}, backoff.NewExponentialBackOff())
		if err != nil {
			return nil, err
		}
	}
	return streamers, nil
}
