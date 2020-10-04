package main

func (b *twitchBot) handleTwitter(_ []string) error {
	b.respond("I can be found on twitter: twitter.com/kinbiko")
	return nil
}
