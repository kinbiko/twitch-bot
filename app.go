package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/kinbiko/bugsnag"
	"github.com/sirupsen/logrus"
)

type twitchBot struct {
	client      *twitch.Client
	channelName string
	*logrus.Logger
	notifier          *bugsnag.Notifier
	env               map[string]string
	handlers          map[string]func(args []string) error
	unpopularOpinions []string
}

func (b *twitchBot) respond(msg string) {
	b.client.Say(b.channelName, msg)
}

// !unpopularopinion
func (b *twitchBot) handleUnpopularOpinion(_ []string) error {
	b.respond(b.unpopularOpinions[rand.Intn(len(b.unpopularOpinions))])
	return nil
}

func (b *twitchBot) setUpHandlers() {
	h := map[string]func(args []string) error{}
	h["!unpopularopinion"] = b.handleUnpopularOpinion
	b.handlers = h
}

func (b *twitchBot) onChatMsg(msg twitch.PrivateMessage) {
	ctx := context.Background()
	// Print the message in the console
	b.Infof("%s: %s\n", msg.User.Name, msg.Message)

	split := strings.Split(msg.Message, " ")
	if err := b.handlers[split[0]](split[1:]); err != nil {
		b.notifier.Notify(ctx, err)
	}
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
