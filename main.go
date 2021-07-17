package main

import (
	"github.com/lewapkon/twitch-clips-poster/pkg/twitchtracker"
	"log"
)

func main() {
	streamers, err := twitchtracker.FindTopStreamers()
	if err != nil {
		log.Panic(err)
		return
	}
	log.Print(streamers)
	log.Print(len(streamers))
}
