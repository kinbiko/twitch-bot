package main

import (
	"math/rand"

	"github.com/gempir/go-twitch-irc/v2"
)

// !unpopularopinion
func (b *twitchBot) handleUnpopularOpinion(_ *twitch.PrivateMessage) error {
	opinions := []string{
		"Consistency is overrated",
		"Ship on Fridays",
		"TDD",
		"Best practices are harmful",
	}
	b.respond(opinions[rand.Intn(len(opinions))])
	return nil
}
