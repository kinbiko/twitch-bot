package main

import (
	"context"
	"strings"
	"sync"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kinbiko/bugsnag"
	"github.com/sirupsen/logrus"
)

type twitchBot struct {
	client      *twitch.Client
	channelName string
	userName    string
	*logrus.Logger
	notifier *bugsnag.Notifier
	env      map[string]string
	handlers map[string]func(msg *twitch.PrivateMessage) error

	xkcdOnce sync.Once
	xkcdData []*xkcdData
}

func (b *twitchBot) respond(msg string) {
	b.Infof("%s: %s", b.userName, msg)
	b.client.Say(b.channelName, msg)
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
	if err := handler(&msg); err != nil {
		b.notifier.Notify(ctx, err)
	}
}
