package main

import (
	"context"
	"strings"

	"github.com/fatih/color"
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
	handlers map[string]func(msg *twitch.PrivateMessage) error

	xkcdData []*xkcdData
}

var (
	blue   = color.New(color.FgBlue, color.Bold).SprintfFunc()
	red    = color.New(color.FgRed, color.Bold).SprintfFunc()
	yellow = color.New(color.FgYellow, color.Bold).SprintfFunc()
)

func (b *twitchBot) respond(msg string) {
	b.Infof("%s: %s", blue(b.userName), msg)
	b.client.Say(b.channelName, msg)
}

func (b *twitchBot) onChatMsg(msg twitch.PrivateMessage) {
	ctx := b.notifier.WithUser(context.Background(), bugsnag.User{Name: msg.User.Name, ID: msg.User.ID})
	ctx = b.notifier.WithMetadatum(ctx, "chat", "message", msg.Message)

	col := yellow
	switch msg.User.Name {
	case "kinbiko":
		col = red
	case "kinbikobot":
		col = blue
	}
	b.Infof("%s: %s\n", col(msg.User.Name), msg.Message)

	split := strings.Split(msg.Message, " ")
	ctx = b.notifier.WithBugsnagContext(ctx, split[0])
	handler, ok := b.handlers[split[0]]
	if !ok {
		return // no handler in this case
	}
	if err := handler(&msg); err != nil {
		b.notifier.Notify(ctx, err)
	}
}
