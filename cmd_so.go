package main

import (
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func (b *twitchBot) handleSo(msg *twitch.PrivateMessage) error {
	if name := msg.User.Name; name != b.channelName {
		return fmt.Errorf("!so can only be invoked by %s, but was invoked by %s", b.channelName, name)
	}
	args := strings.Split(msg.Message, " ")
	if l := len(args); l != 2 {
		return fmt.Errorf("expected 1 argument to !so but got %d: %v", l, args)
	}
	b.respond(fmt.Sprintf("git checkout --streamer twitch.tv/%s", args[1]))
	return nil
}
