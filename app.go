package bot

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
	client *twitch.Client
	*logrus.Logger
	notifier          *bugsnag.Notifier
	env               map[string]string
	handlers          map[string]func(args []string) error
	unpopularOpinions []string
}

func Start() error {
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
		client:   client,
		Logger:   logrus.New(),
		notifier: n,
		env:      env,
		unpopularOpinions: []string{
			"consistency is overrated",
			"ship on Fridays",
			"TDD",
			"best practices are harmful",
		},
	}
	bot.setUpHandlers()

	client.OnPrivateMessage(bot.onChatMsg)
	client.Join(env["CHANNEL_NAME"])
	bot.Info("starting bot...")
	return client.Connect() // this line blocks
}

func (b *twitchBot) respond(msg string) {
	b.client.Say(b.env["CHANNEL_NAME"], msg)
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
