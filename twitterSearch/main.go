package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumer_key      = ""
	consumer_secret   = ""
	accessToken       = ""
	accessTokenSecret = ""
)

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return api
}

func main() {

	api := GetTwitterApi()
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "<Usage> %s word\n", os.Args[0])
		os.Exit(1)
	}
	word := os.Args[1]

	searchResult, _ := api.GetSearch(string(word), nil)

	f, err := os.Create("tweet.txt")
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
	defer f.Close()
	for _, tweet := range searchResult.Statuses {

		io.Copy(f, strings.NewReader(tweet.Text))
	}
}
