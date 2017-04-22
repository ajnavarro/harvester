package harvester

const (
	Email        = MetadataKey("email")
	Company      = MetadataKey("company")
	GithubUser   = MetadataKey("githubUser")
	TwitterUser  = MetadataKey("twitterUser")
	PersonalSite = MetadataKey("personalSite")
	Avatar       = MetadataKey("avatar")
	Name         = MetadataKey("name")
	Location     = MetadataKey("location")
	Description  = MetadataKey("description")
	Language     = MetadataKey("Language")
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
