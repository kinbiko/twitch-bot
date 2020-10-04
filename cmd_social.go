package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handleSocial(_ *twitch.PrivateMessage) error {
	b.respond("Wanna hang out outside the stream? github.com/kinbiko twitter.com/kinbiko https://discord.gg/PCDafQk")
	return nil
}
