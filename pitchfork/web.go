package pitchfork

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ajnavarro/harvester"
)

type Web struct {
}

func NewWeb() *Web {
	return &Web{}
}
func (p *Web) Name() string {
	return "web"
}
func (p *Web) Harvest(d harvester.Data) error {
	metadatas := d[harvester.PersonalSite]
	for _, m := range metadatas {
		doc, err := goquery.NewDocument(m.Value)
		if err != nil {
			return err
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Find("href").Text())
		})
	}

	return nil
}
