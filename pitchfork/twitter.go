package pitchfork

import (
	"github.com/ajnavarro/harvester"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

type Twitter struct {
	c *twitter.Client
}

func NewTwitter(config *TwitterConfig) *Twitter {
	c := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.Token, config.TokenSecret)
	httpClient := c.Client(oauth1.NoContext, token)

	return &Twitter{
		c: twitter.NewClient(httpClient),
	}

}
func (p *Twitter) Name() string {
	return "twitter"
}
func (p *Twitter) Harvest(d harvester.Data) error {
	meta := d[harvester.TwitterUser]

	var twitterUsers []string
	for _, m := range meta {
		twitterUsers = append(twitterUsers, m.Value)
	}

	var users []twitter.User
	var err error
	if len(twitterUsers) != 0 {
		users, _, err = p.c.Users.Lookup(&twitter.UserLookupParams{
			ScreenName: twitterUsers,
		})
	}

	if err != nil {
		return err
	}

	for _, u := range users {
		d.Add(harvester.Description, u.Description, p.Name())
		d.Add(harvester.Email, u.Email, p.Name())
		d.Add(harvester.Language, u.Lang, p.Name())
		d.Add(harvester.Location, u.Location, p.Name())
		d.Add(harvester.Name, u.Name, p.Name())
		d.Add(harvester.Avatar, u.ProfileBannerURL, p.Name())
		d.Add(harvester.Avatar, u.ProfileImageURL, p.Name())
		d.Add(harvester.Timezone, u.Timezone, p.Name())
	}

	return nil
}
