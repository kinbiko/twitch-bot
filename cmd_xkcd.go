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

func (b *twitchBot) setupXKCD(ctx context.Context) error {
	jsonFile, err := os.Open("./xkcdump.json")
	if err != nil {
		return b.notifier.Wrap(ctx, err, "unable to open xkcdump.json")
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return b.notifier.Wrap(ctx, err, "unable to read xkcdump.json file")
	}

	err = json.Unmarshal(byteValue, &b.xkcdData)
	if err != nil {
		return b.notifier.Wrap(ctx, err, "unable to read parse xkcdump.json as JSON")
	}
	return nil
}

func (b *twitchBot) handleXKCD(msg *twitch.PrivateMessage) error {
	args := strings.Split(msg.Message, " ")
	if len(args) < 2 {
		b.respond("There's no XKCD for that: https://xkcd.com/404")
		return nil
	}
	arg := strings.ToLower(args[1])

	scores := map[int]int{}
	for _, comic := range b.xkcdData {
		// title has a weight of 10
		// Transcript has a weight of 3
		// alt has a weight of 1
		wordsTitle, wordsTranscript, wordsAlt := strings.Split(comic.SafeTitle, " "), strings.Split(comic.Transcript, " "), strings.Split(comic.Alt, " ")
		for _, word := range wordsTitle {
			if strings.ToLower(word) == arg {
				scores[comic.Num] += 10
			}
		}
		for _, word := range wordsTranscript {
			if strings.ToLower(word) == arg {
				scores[comic.Num] += 3
			}
		}
		for _, word := range wordsAlt {
			if strings.ToLower(word) == arg {
				scores[comic.Num] += 1
			}
		}
	}

	mostRelevantComic, maxScore := 404, 0
	for comicNum, score := range scores {
		if score > maxScore {
			maxScore, mostRelevantComic = score, comicNum
		}
	}

	if mostRelevantComic == 404 {
		b.respond("There isn't an XKCD for that: https://xkcd.com/404")
		return nil
	}

	b.respond("There's an xkcd for that: " + fmt.Sprintf(linkFormat, mostRelevantComic))
	return nil
}
