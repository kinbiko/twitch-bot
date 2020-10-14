package main

import (
	"fmt"
	"math/rand"

	"github.com/gempir/go-twitch-irc/v2"
)

func (b *twitchBot) handleLurk(msg *twitch.PrivateMessage) error {
	un := msg.User.Name
	lurkResponses := []string{
		// Taken from Abed's Batman speech from Community.
		fmt.Sprintf("If %s chats, there can be no party. %s must be out there in the night, staying vigilant.", un, un),
		fmt.Sprintf("Wherever a party needs to be saved, %s is there.", un),
		fmt.Sprintf("Wherever there are masks, wherever there’s tomfoolery and joy, %s is there.", un),
		fmt.Sprintf("But sometimes %s is not cause they're out in the night, staying vigilant.", un),
		fmt.Sprintf("Watching. !lurk-ing. Running. Jumping. Hurtling. Sleeping. No, %s can’t sleep. You sleep. %s is awake. %s doesn't sleep.", un, un, un),
		fmt.Sprintf("%s doesn't blink. Is %s a bird? No. %s is a bat. %s is Batman. Or are they? Yes, they're Batman.", un, un, un, un),
	}
	b.respond(lurkResponses[rand.Intn(len(lurkResponses))])
	return nil
}
