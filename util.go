package oxpecker

import (
	"github.com/dghubble/go-twitter/twitter"
	"strings"
)

func ReplaceTweetURLs(tweet *twitter.Tweet) *twitter.Tweet {
	tweet.FullText = replaceStringURLs(tweet.FullText, tweet.Entities.Urls)
	tweet.Text = replaceStringURLs(tweet.Text, tweet.Entities.Urls)
	return tweet
}

func replaceStringURLs(str string, urls []twitter.URLEntity) string {
	for _, url := range urls {
		str = strings.Replace(str, url.URL, url.ExpandedURL, -1)
	}

	return str
}
