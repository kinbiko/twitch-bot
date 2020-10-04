package main

import "github.com/gempir/go-twitch-irc/v2"

func (b *twitchBot) handleDotfiles(_ *twitch.PrivateMessage) error {
	b.respond("My vim/zsh/tmux/terminal/keyboard config can be found on github.com/kinbiko/dotfiles")
	return nil
}
