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
		str = strings.ReplaceAll(str, url.URL, url.ExpandedURL)
	}

	return str
}

func RemoveMediaURLs(tweet *twitter.Tweet) *twitter.Tweet {
	tweet.FullText = removeMediaUrlsString(tweet.FullText, tweet.Entities.Media)
	tweet.Text = removeMediaUrlsString(tweet.Text, tweet.Entities.Media)
	return tweet
}

func removeMediaUrlsString(str string, media []twitter.MediaEntity) string {
	for _, medium := range media {
		str = strings.ReplaceAll(str, "", medium.URL)
	}

	return str
}
