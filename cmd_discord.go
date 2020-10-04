package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handleDiscord(_ *twitch.PrivateMessage) error {
	b.respond("Join the discord server: https://discord.gg/PCDafQk")
	return nil
}
