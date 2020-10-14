package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) setUpRoutes() {
	b.handlers = map[string]func(msg *twitch.PrivateMessage) error{
		// Simple responses
		"!discord":     b.alwaysRespond("Join the discord server: https://discord.gg/PCDafQk"),
		"!dotfiles":    b.alwaysRespond("My vim/zsh/tmux/terminal/keyboard config can be found on github.com/kinbiko/dotfiles"),
		"!github":      b.alwaysRespond("Follow my coding activities on GitHub: github.com/kinbiko"),
		"!phonenumber": b.alwaysRespond("My phone number is 0118 999 881 999 119 725... 3"),
		"!social":      b.alwaysRespond("Wanna hang out outside the stream? github.com/kinbiko twitter.com/kinbiko https://discord.gg/PCDafQk"),
		"!twitter":     b.alwaysRespond("I can be found on twitter: twitter.com/kinbiko"),

		// Commands with a bit more meat
		"!commands":         b.handleCommands,
		"!lurk":             b.handleLurk,
		"!so":               b.handleSo,
		"!unpopularopinion": b.handleUnpopularOpinion,
		"!xkcd":             b.handleXKCD,
	}
}

func (b *twitchBot) alwaysRespond(response string) func(msg *twitch.PrivateMessage) error {
	return func(_ *twitch.PrivateMessage) error {
		b.respond(response)
		return nil
	}
}
