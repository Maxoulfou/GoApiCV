package routes

import (
	"net/http"

	"github.com/Devansh3712/portfolio/api"
	"github.com/Devansh3712/portfolio/models"
	"github.com/gin-gonic/gin"
)

// @Summary		Get miscellaneous data of user
// @Produce 	json
// @Success 	200		{object} 	map[string]any
// @Router 		/about [get]
func Misc(ctx *gin.Context) {
	var recentTweet models.Tweet
	var latestCommit *models.Commit
	tweets, _ := api.GetRecentTweets()
	for _, tweet := range tweets.Data {
		if tweet.InReplyTo == "" {
			recentTweet.ID = tweet.ID
			recentTweet.Text = tweet.Text
			recentTweet.CreatedAt = tweet.CreatedAt
			break
		}
	}
	latestCommit, _ = api.GetLatestCommit()
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"recent_tweet":  recentTweet,
		"latest_commit": latestCommit,
	})
}
