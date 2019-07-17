package oxpecker

import (
	"errors"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/McKael/madon"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const appName = "Oxpecker"

type Config struct {
	TwitterAccounts []struct {
		ConsumerKey    string
		ConsumerSecret string
		AccessToken    string
		AccessSecret   string
		TrackUsers     []string
	}

	MastodonAccounts []struct {
		InstanceName string
		AppID        string
		AppSecret    string
		UserToken    string
	}
}

type Oxpecker struct {
	birds []struct {
		client *twitter.Client
		track  []string
	}
	elephants []*madon.Client
}

func New(c *Config) (*Oxpecker, error) {
	ox := &Oxpecker{}

	for _, twtconfig := range c.TwitterAccounts {
		twitterConfig := oauth1.NewConfig(twtconfig.ConsumerKey, twtconfig.ConsumerSecret)
		twitterToken := oauth1.NewToken(twtconfig.AccessToken, twtconfig.AccessSecret)

		client := twitter.NewClient(twitterConfig.Client(oauth1.NoContext, twitterToken))
		_, _, err := client.Accounts.VerifyCredentials(nil)
		if err != nil {
			log.Print("error building twitter client: " + err.Error())
		} else {
			ox.birds = append(ox.birds, struct {
				client *twitter.Client
				track  []string
			}{client: client, track: twtconfig.TrackUsers})
		}
	}

	for _, mstdconfig := range c.MastodonAccounts {
		token := madon.UserToken{}
		token.AccessToken = mstdconfig.UserToken

		client, err := madon.RestoreApp(appName, mstdconfig.InstanceName, mstdconfig.AppID, mstdconfig.AppSecret, &token)
		if err != nil {
			log.Print("error building mastodon client: " + err.Error())
			continue
		}

		_, err = client.GetCurrentAccount()
		if err != nil {
			log.Print("error building mastodon client: " + err.Error())
			continue
		}
		ox.elephants = append(ox.elephants, client)
	}

	if len(ox.elephants) == 0 {
		return nil, errors.New("at least one mastodon account is required, but none could be instantiated successfully")
	}

	if len(ox.birds) == 0 {
		return nil, errors.New("at least one twitter account is required, but none could be instantiated successfully")
	}

	return ox, nil
}

func NewFromTomlFile(path string) (*Oxpecker, error) {
	config := &Config{}
	_, err := toml.DecodeFile(path, config)

	if err != nil {
		return nil, errors.New("error parsing config file: " + err.Error())
	}

	return New(config)
}

func (ox *Oxpecker) Run() error {
	multichan := make(chan *twitter.Tweet, 4)

	for _, bird := range ox.birds {
		users, _, err := bird.client.Users.Lookup(&twitter.UserLookupParams{
			ScreenName: bird.track,
		})
		if err != nil {
			log.Println("error getting own id from twitter: " + err.Error())
			continue
		}

		filterparams := twitter.StreamFilterParams{}

		from := ""
		for _, user := range users {
			filterparams.Follow = append(filterparams.Follow, user.IDStr)
			from += ", @" + user.ScreenName
		}

		stream, err := bird.client.Streams.Filter(&filterparams)
		if err != nil {
			log.Println("error streaming from twitter: " + err.Error())
			continue
		}

		go func() {
			log.Printf("Listening for tweets from " + strings.TrimPrefix(from, ", "))
			for event := range stream.Messages {
				tweet, ok := event.(*twitter.Tweet)
				if !ok {
					log.Println("Got event but it's not a tweet, skipping")
					continue
				}

				if (strings.HasPrefix(tweet.Text, "@") || tweet.InReplyToStatusIDStr != "") && tweet.InReplyToUserIDStr != tweet.User.IDStr {
					log.Printf("Tweet %s is a reply, skipping", tweet.IDStr)
					continue
				}

				if tweet.QuotedStatusID != 0 {
					log.Printf("Tweet %s is a quote, skipping", tweet.IDStr)
					continue
				}

				if tweet.RetweetedStatus != nil {
					continue
				}

				multichan <- tweet
			}
		}()
	}

	for tweet := range multichan {
		text := tweet.Text
		if tweet.FullText != "" {
			text = tweet.FullText
		}

		for _, elephant := range ox.elephants {
			_, err := elephant.PostStatus(madon.PostStatusParams{
				Text:       text,
				Visibility: "public",
			})

			if err != nil {
				log.Printf("error posting tweet %s to mastodon: %s", tweet.IDStr, err.Error())
			}
		}
	}

	return nil
}
