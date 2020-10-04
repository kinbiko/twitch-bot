package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type xkcdData struct {
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
}

func main() {
	data := make([]*xkcdData, 2367)
	for i := 1; i <= 2367; i++ {
		apiEndpoint := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		resp, err := http.Get(apiEndpoint)
		if err != nil {
			panic(err)
		}
		val := &xkcdData{}
		err = json.NewDecoder(resp.Body).Decode(val)
		if err != nil {
			panic(err)
		}
		data[i-1] = val
		file, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("xkcdump.json", file, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("completed %d comic metadata grabs\n", i)
		time.Sleep(1 * time.Second)
	}
}
