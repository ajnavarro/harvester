package pitchfork

import (
	"net/http"
	"time"

	"encoding/json"
	"fmt"
	"github.com/ajnavarro/harvester"
)

const (
	httpTimeout = 30 * time.Second

	baseUrl = "https://keybase.io/_/api/1.0/user/lookup.json?%s&fields=basics,pictures,profile,proofs_summary"
)

var searchParams map[string]harvester.MetadataKey = map[string]harvester.MetadataKey{
	"github=%s":     harvester.GithubUser,
	"domain=%s":     harvester.PersonalSite,
	"twitter=%s":    harvester.TwitterUser,
	"reddit=%s":     harvester.RedditUser,
	"hackernews=%s": harvester.HackernewsUser,
}

type Keybase struct {
	c *http.Client
}

func NewKeybase() harvester.Pitchfork {
	c := &http.Client{}
	c.Timeout = httpTimeout

	return &Keybase{c}
}

func (p *Keybase) Harvest(d harvester.Data) error {
	for param, key := range searchParams {
		if err := p.fetchData(d, key, param); err != nil {
			return err
		}
	}

	return nil
}

func (p *Keybase) fetchData(d harvester.Data, k harvester.MetadataKey, queryParam string) error {
	return d.ForEach(k, func(m *harvester.Metadata) error {
		res, err := p.c.Get(fmt.Sprintf(baseUrl, fmt.Sprintf(queryParam, m.Value)))
		if err != nil {
			return err
		}

		defer res.Body.Close()
		if res.StatusCode >= 400 {
			return fmt.Errorf("request error. Status code %s, %s", res.Status, res.Status)
		}

		var record *UserModel
		if err := json.NewDecoder(res.Body).Decode(&record); err != nil {
			return err
		}

		for _, t := range record.Them {
			d.Add(harvester.KeybaseUser, t.Basics.Username, p.Name())
			d.Add(harvester.Description, t.Profile.Bio, p.Name())
			d.Add(harvester.Name, t.Profile.FullName, p.Name())
			d.Add(harvester.Location, t.Profile.Location, p.Name())

			for _, f := range t.ProofsSummary.ByProofType.Facebook {
				d.Add(harvester.FacebookUser, f.Nametag, p.Name())

			}
			for _, gs := range t.ProofsSummary.ByProofType.GenericWebSite {
				d.Add(harvester.PersonalSite, gs.ServiceURL, p.Name())
			}

			for _, gh := range t.ProofsSummary.ByProofType.Github {
				d.Add(harvester.GithubUser, gh.Nametag, p.Name())
			}

			for _, hn := range t.ProofsSummary.ByProofType.Hackernews {
				d.Add(harvester.HackernewsUser, hn.Nametag, p.Name())
			}
			for _, r := range t.ProofsSummary.ByProofType.Reddit {
				d.Add(harvester.RedditUser, r.Nametag, p.Name())
			}

			for _, t := range t.ProofsSummary.ByProofType.Twitter {
				d.Add(harvester.TwitterUser, t.Nametag, p.Name())
			}

			for _, ca := range t.CryptocurrencyAddresses.Bitcoin {
				d.Add(harvester.BitcoinAddress, ca.Address, p.Name())
			}
		}

		return nil
	})

}

func (p *Keybase) Name() string {
	return "keybase"
}

type UserModel struct {
	Status struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
	Them []struct {
		ID     string `json:"id"`
		Basics struct {
			Username      string `json:"username"`
			Ctime         int    `json:"ctime"`
			Mtime         int    `json:"mtime"`
			IDVersion     int    `json:"id_version"`
			TrackVersion  int    `json:"track_version"`
			LastIDChange  int    `json:"last_id_change"`
			UsernameCased string `json:"username_cased"`
		} `json:"basics"`
		Profile struct {
			Mtime    int    `json:"mtime"`
			FullName string `json:"full_name"`
			Location string `json:"location"`
			Bio      string `json:"bio"`
		} `json:"profile"`
		CryptocurrencyAddresses struct {
			Bitcoin []struct {
				Address string `json:"address"`
				SigID   string `json:"sig_id"`
			} `json:"bitcoin"`
		} `json:"cryptocurrency_addresses"`
		ProofsSummary struct {
			ByProofType struct {
				Twitter []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"twitter"`
				Github []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"github"`
				Reddit []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"reddit"`
				Hackernews []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"hackernews"`
				Facebook []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"facebook"`
				GenericWebSite []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"generic_web_site"`
				DNS []struct {
					ProofType         string `json:"proof_type"`
					Nametag           string `json:"nametag"`
					State             int    `json:"state"`
					ProofURL          string `json:"proof_url"`
					SigID             string `json:"sig_id"`
					ProofID           string `json:"proof_id"`
					HumanURL          string `json:"human_url"`
					ServiceURL        string `json:"service_url"`
					PresentationGroup string `json:"presentation_group"`
					PresentationTag   string `json:"presentation_tag"`
				} `json:"dns"`
			} `json:"by_proof_type"`
		} `json:"proofs_summary"`
		Pictures struct {
			Primary struct {
				URL    string      `json:"url"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Source interface{} `json:"source"`
			} `json:"primary"`
		} `json:"pictures"`
	} `json:"them"`
}
