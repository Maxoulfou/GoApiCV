package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Devansh3712/portfolio/models"
)

func GetRecentTweets() (*models.Twitter, error) {
	api := "https://api.twitter.com/2/tweets/search/recent?query=from:" +
		os.Getenv("TWITTER_ID") +
		"&tweet.fields=created_at,in_reply_to_user_id"
	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_TOKEN"))
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var tweetData models.Twitter
	if err := json.Unmarshal(data, &tweetData); err != nil {
		return nil, err
	}
	return &tweetData, nil
}
