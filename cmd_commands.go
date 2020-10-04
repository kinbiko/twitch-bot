package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func (b *twitchBot) handleCommands(msg *twitch.PrivateMessage) error {
	cmds := []string{}
	for cmd := range b.handlers {
		cmds = append(cmds, cmd)
	}
	sort.Strings(cmds)
	res := fmt.Sprintf("All commands available are: %s", strings.Join(cmds, ", "))
	b.respond(res)
	return nil
}
