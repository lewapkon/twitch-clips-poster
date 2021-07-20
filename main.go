package main

import (
	"github.com/lewapkon/twitch-clips-poster/pkg/config"
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
	log.Print(streamers)
	log.Print(len(streamers))
}
