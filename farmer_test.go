package harvester_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ajnavarro/harvester"
	"github.com/ajnavarro/harvester/pitchfork"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	tc := &pitchfork.TwitterConfig{
		ConsumerKey:    "",
		ConsumerSecret: "",
		Token:          "",
		TokenSecret:    "",
	}

	farmer := harvester.NewFarmer([]harvester.Pitchfork{
		pitchfork.NewGithub(os.Getenv("GITHUB_TOKEN")),
		pitchfork.NewTwitter(tc),
		pitchfork.NewWeb(),
	})

	data, err := farmer.Farm(harvester.Seeds{
		harvester.GithubUser:  "some",
	})

	assert.NoError(t, err)

	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "    ")
	err = enc.Encode(data)
	assert.NoError(t, err)
	fmt.Println(buf.String())
}
