package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

type xkcdData struct {
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

const linkFormat = "https://xkcd.com/%d/"

func (b *twitchBot) handleXkcd(msg *twitch.PrivateMessage) error {
	var comic xkcdData
	b.respond("There's an xkcd for that: " + fmt.Sprintf(linkFormat, comic.Num))
	return nil
}
