package pitchfork

import (
	"context"
	"net/http"

	"github.com/ajnavarro/harvester"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	ctx context.Context
	c   *github.Client
}

func NewGithub(token string) harvester.Pitchfork {
	c := http.DefaultClient
	ctx := context.Background()
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		c = oauth2.NewClient(ctx, ts)
	}

	return &Github{
		ctx: ctx,
		c:   github.NewClient(c),
	}

}
func (p *Github) Name() string {
	return "github"
}
func (p *Github) Harvest(d harvester.Data) error {
	return d.ForEach(harvester.GithubUser, func(m *harvester.Metadata) error {
		user, _, err := p.c.Users.Get(p.ctx, m.Value)
		if err != nil {
			return err
		}

		d.Add(harvester.Name, dereference(user.Name), p.Name())
		d.Add(harvester.Email, dereference(user.Email), p.Name())
		d.Add(harvester.Avatar, dereference(user.AvatarURL), p.Name())
		d.Add(harvester.PersonalSite, dereference(user.Blog), p.Name())
		d.Add(harvester.Company, dereference(user.Company), p.Name())
		d.Add(harvester.Location, dereference(user.Location), p.Name())
		d.Add(harvester.GithubUser, dereference(user.Login), p.Name())

		return nil
	})
}

func dereference(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
