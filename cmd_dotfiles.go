package main

func (b *twitchBot) handleDotfiles(_ []string) error {
	b.respond("My vim/zsh/tmux/terminal/keyboard config can be found on github.com/kinbiko/dotfiles")
	return nil
}
