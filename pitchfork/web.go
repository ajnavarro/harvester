package pitchfork

import (
	"net/url"
	"path"

	"github.com/ajnavarro/harvester"

	"github.com/PuerkitoBio/goquery"
	log "github.com/inconshreveable/log15"
)

type Web struct {
}

func NewWeb() harvester.Pitchfork {
	return &Web{}
}
func (p *Web) Name() string {
	return "web"
}
func (p *Web) Harvest(d harvester.Data) error {
	meta := d[harvester.PersonalSite]
	for _, m := range meta {
		doc, err := goquery.NewDocument(m.Value)
		if err != nil {
			return err
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			v, ok := s.Attr("href")
			if ok {
				u, err := url.Parse(v)
				if err != nil {
					log.Error("Error parsing url", "url", v, "err", err)
				}

				// TODO improve url user detection
				switch u.Host {
				case "github.com", "www.github.com":
					d.Add(harvester.GithubUser, path.Base(u.Path), p.Name())
				case "linkedin.com", "www.linkedin.com":
					if path.Base(u.Path) == "in" {
						_, user := path.Split(u.Path)
						d.Add(harvester.LinkedinUser, user, p.Name())
					}
				case "twitter.com", "www.twitter.com":
					d.Add(harvester.TwitterUser, path.Base(u.Path), p.Name())
				default:
					log.Debug("URL not associated", "URL", u)
				}
			}

		})
	}

	return nil
}
