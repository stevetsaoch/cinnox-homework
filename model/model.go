package model

// Config structure
type Config struct {
	LineChannelSecret      string `mapstructure:"lineChannelSecret"`
	LineChannelAccessToken string `mapstructure:"lineChannelAccessToken"`
	MongoUser              string `mapstructure:"mongodbUser"`
	MongoPwd               string `mapstructure:"mongodbPwd"`
	DatabaseName           string `mapstructure:"databaseName"`
	CollectionName         string `mapstructure:"collectionName"`
}
