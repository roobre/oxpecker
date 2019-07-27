package oxpecker

import (
	"github.com/McKael/madon"
	"github.com/dghubble/go-twitter/twitter"
)

type Elephant struct {
	*madon.Client
	postMap map[int64]int64
}

func (e *Elephant) Mirror(tweet *twitter.Tweet) error {
	tweet = ReplaceTweetURLs(RemoveMediaURLs(tweet))

	text := tweet.Text
	if tweet.Truncated {
		text = tweet.ExtendedTweet.FullText
	}

	status, err := e.PostStatus(madon.PostStatusParams{
		Text:       text,
		Visibility: "public",
		InReplyTo:  e.postMap[tweet.InReplyToStatusID],
	})

	if err != nil {
		return err
	}

	e.postMap[tweet.ID] = status.ID
	return nil
}
