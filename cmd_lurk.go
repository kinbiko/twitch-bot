package main

import (
	"math/rand"

	"github.com/gempir/go-twitch-irc/v2"
)

func (b *twitchBot) handleLurk(msg *twitch.PrivateMessage) error {
	lurkResponses := []string{
		// Taken from Abed's Batman speech from Community.
		"If %s chats, there can be no party. %s must be out there in the night, staying vigilant.",
		"Wherever a party needs to be saved, %s is there.",
		"Wherever there are masks, wherever there’s tomfoolery and joy, %s is there.",
		"But sometimes %s is not cause they're out in the night, staying vigilant.",
		"Watching. !lurk-ing. Running. Jumping. Hurtling. Sleeping. No, %s can’t sleep. You sleep. %s is awake. %s doesn't sleep.",
		"%s doesn't blink. Is %s a bird? No. %s is a bat. %s is Batman. Or are they? Yes, they're Batman.`)",
	}
	b.respond(lurkResponses[rand.Intn(len(lurkResponses))])
	return nil
}
