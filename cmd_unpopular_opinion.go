package main

import "math/rand"

func unpopularOpinions() []string {
	return []string{
		"Consistency is overrated",
		"Ship on Fridays",
		"TDD",
		"Best practices are harmful",
	}
}

// !unpopularopinion
func (b *twitchBot) handleUnpopularOpinion(_ []string) error {
	b.respond(b.unpopularOpinions[rand.Intn(len(b.unpopularOpinions))])
	return nil
}
