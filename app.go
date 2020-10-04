package main

import (
	"context"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
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

func (b *twitchBot) setUpHandlers() {
	h := map[string]func(args []string) error{}
	h["!unpopularopinion"] = b.handleUnpopularOpinion
	h["!dotfiles"] = b.handleDotfiles
	h["!twitter"] = b.handleTwitter
	h["!discord"] = b.handleDiscord
	h["!github"] = b.handleGitHub
	h["!social"] = b.handleSocial
	b.handlers = h
}

func (b *twitchBot) onChatMsg(msg twitch.PrivateMessage) {
	ctx := b.notifier.WithUser(context.Background(), bugsnag.User{Name: msg.User.Name, ID: msg.User.ID})
	// Print the message in the console
	b.Infof("%s: %s\n", msg.User.Name, msg.Message)

	split := strings.Split(msg.Message, " ")
	handler, ok := b.handlers[split[0]]
	if !ok {
		return // no handler in this case
	}
	if err := handler(split[1:]); err != nil {
		b.notifier.Notify(ctx, err)
	}
}
