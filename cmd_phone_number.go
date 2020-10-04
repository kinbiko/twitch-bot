package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handlePhoneNumber(msg *twitch.PrivateMessage) error {
	b.respond("0118 999 881 999 119 725... 3")
	return nil
}
