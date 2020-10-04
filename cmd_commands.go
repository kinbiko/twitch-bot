package main

import (
	"fmt"
	"sort"
	"strings"
)

func (b *twitchBot) handleCommands(_ []string) error {
	cmds := []string{}
	for cmd := range b.handlers {
		cmds = append(cmds, cmd)
	}
	sort.Strings(cmds)
	msg := fmt.Sprintf("All commands available are: %s", strings.Join(cmds, ", "))
	b.respond(msg)
	return nil
}
