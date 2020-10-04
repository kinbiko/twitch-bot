package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handleGitHub(_ *twitch.PrivateMessage) error {
	b.respond("Follow my coding activities on GitHub: github.com/kinbiko")
	return nil

}
