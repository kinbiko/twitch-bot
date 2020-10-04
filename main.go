package main

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/kinbiko/bugsnag"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func start() error {
	env, err := readEnv()
	if err != nil {
		return err
	}

	client := twitch.NewClient(env["BOT_USERNAME"], env["OAUTH_TOKEN"])

	n, err := bugsnag.New(bugsnag.Configuration{
		APIKey:       env["BUGSNAG_API_KEY"],
		AppVersion:   "0.0.1-dev",
		ReleaseStage: "dev",
	})

	if err != nil {
		return err
	}

	bot := &twitchBot{
		client:            client,
		channelName:       env["CHANNEL_NAME"],
		Logger:            logrus.New(),
		notifier:          n,
		unpopularOpinions: unpopularOpinions(),
	}
	bot.setUpHandlers()

	client.OnPrivateMessage(bot.onChatMsg)
	client.Join(env["CHANNEL_NAME"])
	bot.Info("starting bot...")
	return client.Connect() // this line blocks
}

func readEnv() (map[string]string, error) {
	env, err := godotenv.Read()
	if err != nil {
		return nil, fmt.Errorf("unable to load environment from .env file: %w", err)
	}
	for _, s := range []string{
		"BOT_USERNAME",
		"BUGSNAG_API_KEY",
		"CHANNEL_NAME",
		"OAUTH_TOKEN",
	} {
		if env[s] == "" {
			return nil, fmt.Errorf("couldn't find '%s' in .env file", s)
		}
	}
	return env, nil
}
