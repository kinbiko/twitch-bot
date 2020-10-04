package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
	var (
		ctx = context.Background()
		err error
	)

	xkcdDump := []xkcdData{}
	b.xkcdOnce.Do(func() {
		jsonFile, err := os.Open("./xkcdump.json")
		if err != nil {
			err = b.notifier.Wrap(ctx, err, "unable to open xkcdump.json")
			return
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			err = b.notifier.Wrap(ctx, err, "unable to read xkcdump.json file")
			return
		}

		err = json.Unmarshal(byteValue, &xkcdDump)
		if err != nil {
			err = b.notifier.Wrap(ctx, err, "unable to read parse xkcdump.json as JSON")
			return
		}
	})
	if err != nil {
		return err
	}
	comic := xkcdDump[406]
	b.respond("There's an xkcd for that: " + fmt.Sprintf(linkFormat, comic.Num))
	return nil
}
