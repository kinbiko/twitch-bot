package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handleTwitter(_ *twitch.PrivateMessage) error {
	b.respond("I can be found on twitter: twitter.com/kinbiko")
	return nil
}
