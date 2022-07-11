package models

import (
	"encoding/json"
	"os"
	"time"
)

type Tweet struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Twitter struct {
	Data []struct {
		Tweet
		InReplyTo string `json:"in_reply_to_user_id"`
	} `json:"data"`
}

func (data Tweet) MarshalJSON() ([]byte, error) {
	res := map[string]any{
		"url":       "https://twitter.com/" + os.Getenv("TWITTER_ID") + "/status/" + data.ID,
		"tweet":     data.Text,
		"timestamp": data.CreatedAt,
	}
	return json.Marshal(res)
}
