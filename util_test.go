package oxpecker

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"strings"
	"testing"
)

const tco = "://t.co"

func TestReplaceTweetURLs(t *testing.T) {
	tweet := &twitter.Tweet{}
	err := json.Unmarshal([]byte(sampleTweet), tweet)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(tweet.Text, tco) {
		t.Fatal("Original tweet does not contain url")
	}

	tweet = ReplaceTweetURLs(tweet)

	if strings.Contains(tweet.Text, tco) {
		t.Fatal("Replaced tweet contains short url")
	}
	if !strings.Contains(tweet.Text, "cards.twitter.com") {
		t.Fatal("Replaced tweet does not contain full url")
	}
}

const sampleTweet = `{
  "created_at": "Thu Apr 06 15:24:15 +0000 2017",
  "id_str": "850006245121695744",
  "text": "1\/ Today we\u2019re sharing our vision for the future of the Twitter API platform!\nhttps:\/\/t.co\/XweGngmxlP",
  "user": {
    "id": 2244994945,
    "name": "Twitter Dev",
    "screen_name": "TwitterDev",
    "location": "Internet",
    "url": "https:\/\/dev.twitter.com\/",
    "description": "Your official source for Twitter Platform news, updates & events. Need technical help? Visit https:\/\/twittercommunity.com\/ \u2328\ufe0f #TapIntoTwitter"
  },
  "place": {   
  },
  "entities": {
    "hashtags": [      
    ],
    "urls": [
      {
        "url": "https:\/\/t.co\/XweGngmxlP",
        "expanded_url": "https:\/\/cards.twitter.com\/cards\/18ce53wgo4h\/3xo1c",
        "unwound": {
          "url": "https:\/\/cards.twitter.com\/cards\/18ce53wgo4h\/3xo1c",
          "title": "Building the Future of the Twitter API Platform"
        }
      }
    ],
    "user_mentions": [     
    ]
  }
}`
