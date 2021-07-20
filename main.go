package main

import (
	"github.com/lewapkon/twitch-clips-poster/pkg/config"
	"github.com/lewapkon/twitch-clips-poster/pkg/twitch"
	"github.com/lewapkon/twitch-clips-poster/pkg/twitchtracker"
	"log"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}
	streamers, err := twitchtracker.FindTopStreamers(conf.TwitchTracker.Language, conf.TwitchTracker.PagesToLoad)
	if err != nil {
		log.Fatal("Error finding top streamers: ", err)
	}
	twitchClient, err := twitch.NewTwitchClient(conf.Twitch.ClientId, conf.Twitch.ClientSecret)
	if err != nil {
		log.Fatal("Error creating twitch client: ", err)
	}
	users, err := twitchClient.FindUsers(streamers)
	if err != nil {
		log.Fatal("Error finding users: ", err)
	}

	// Display streamer name to user id mapping to test
	for _, user := range users {
		log.Printf("%s: %s", user.Login, user.ID)
	}
}
