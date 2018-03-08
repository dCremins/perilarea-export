package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
)

func Handler(ctx context.Context) (string, error) {
	token := os.Getenv("OAUTH_TOKEN")
	channel := os.Getenv("GEN_CHANNEL")
	client := slack.New(token)

	// Get history
	hist, err := client.GetChannelHistory(channel, slack.NewHistoryParameters())
	if err != nil {
		return "Failed to fetch Channel History", err
	}

	msg := fmt.Sprintf("Archiving %d chat entries", len(hist.Messages))

	if _, _, err := client.PostMessage(channel, msg, slack.NewPostMessageParameters()); err != nil {
		return "Failed to post message to channel", err
	}

	return "Done", nil
}

func main() {
	lambda.Start(Handler)
}
