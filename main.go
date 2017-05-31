package main

import (
	"flag"
	"log"

	"github.com/nlopes/slack"
)

var (
	token   = flag.String("token", "", "Authentication token.")
	channel = flag.String("channel", "", "Slack channel")
	message = flag.String("message", "", "Message to post")
)

func main() {
	flag.Parse()

	if len(*token) == 0 {
		log.Fatal("No access token provided")
	}

	if len(*channel) == 0 {
		log.Fatal("No channel name provided")
	}

	if len(*message) == 0 {
		log.Fatal("No message provided")
	}

	api := slack.New(*token)
	params := slack.PostMessageParameters{
		AsUser: true,
	}
	channelID, timestamp, err := api.PostMessage(*channel, *message, params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
