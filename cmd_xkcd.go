package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	args := strings.Split(msg.Message, " ")
	if len(args) < 2 {
		b.respond("There isn't an XKCD for that: https://xkcd.com/404")
	}
	arg := args[1]

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

	scores := map[int]int{}
	for _, comic := range xkcdDump {
		// title has a weight of 10
		// Transcript has a weight of 3
		// alt has a weight of 1
		wordsTitle, wordsTranscript, wordsAlt := strings.Split(comic.SafeTitle, " "), strings.Split(comic.Transcript, " "), strings.Split(comic.Alt, " ")
		for _, word := range wordsTitle {
			if word == arg {
				scores[comic.Num] += 10
			}
		}
		for _, word := range wordsTranscript {
			if word == arg {
				scores[comic.Num] += 3
			}
		}
		for _, word := range wordsAlt {
			if word == arg {
				scores[comic.Num] += 1
			}
		}
	}

	max := 404
	currentMaxScore := 0
	for num, score := range scores {
		if score > currentMaxScore {
			currentMaxScore = score
			max = num
		}
	}
	if max == 404 {
		b.respond("There isn't an XKCD for that: https://xkcd.com/404")
		return nil
	}

	b.respond("There's an xkcd for that: " + fmt.Sprintf(linkFormat, max))
	return nil
}
