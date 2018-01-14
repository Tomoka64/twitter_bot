package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

const ()

func main() {
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	// searchResult, _ := api.GetSearch("golang", nil)
	// for _, tweet := range searchResult.Statuses {
	// 	fmt.Println(tweet.Text)
	// }

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#love"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Fatal("recerived unexpected value")
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		fmt.Printf("%v\n", t.Text)
		//_, err := api.Retweet(t.Id, false)
		// if err != nil {
		// 	log.Fatal("unable to retweet")
		// 	continue
		// }
		fmt.Printf("%v\n", t.Id)
	}
}
