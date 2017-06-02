package harvester

const (
	Email          = MetadataKey("email")
	Company        = MetadataKey("company")
	GithubUser     = MetadataKey("githubUser")
	TwitterUser    = MetadataKey("twitterUser")
	PersonalSite   = MetadataKey("personalSite")
	Avatar         = MetadataKey("avatar")
	Name           = MetadataKey("name")
	Location       = MetadataKey("location")
	Timezone       = MetadataKey("timezone")
	Description    = MetadataKey("description")
	Language       = MetadataKey("language")
	KeybaseUser    = MetadataKey("keybaseUser")
	FacebookUser   = MetadataKey("facebookUser")
	HackernewsUser = MetadataKey("hackernewsUser")
	RedditUser     = MetadataKey("redditUser")
	LinkedinUser   = MetadataKey("linkedinUser")
	BitcoinAddress = MetadataKey("bitcoinAddress")
)

type Pitchfork interface {
	Harvest(d Data) error
	Name() string
}

type MetadataKey string

type Metadata struct {
	Value     string `json:"value"`
	Pitchfork string `json:"pitchfork"`
	// TODO add more data like percentage of reliability
}
