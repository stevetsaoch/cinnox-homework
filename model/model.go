package model

import "time"

// Config structure
type Config struct {
	LineChannelSecret      string `mapstructure:"lineChannelSecret"`
	LineChannelAccessToken string `mapstructure:"lineChannelAccessToken"`
	MongoUser              string `mapstructure:"mongodbUser"`
	MongoPwd               string `mapstructure:"mongodbPwd"`
	DatabaseName           string `mapstructure:"databaseName"`
	CollectionName         string `mapstructure:"collectionName"`
}

// linebot relate structure
type LineEvent struct {
	Id                string    `json:"id" bson:"_id"`
	Replytoken        string    `json:"replytoken" bson:"replytoken"`
	Type              string    `json:"type" bson:"type"`
	Mode              string    `json:"mode" bson:"mode"`
	Timestamp         time.Time `json:"timestamp" bson:"timestamp"`
	Source            Source    `json:"source" bson:"source"`
	Message           Message   `json:"message" bson:"message"`
	Joined            string    `json:"joined" bson:"joined"`
	Left              string    `json:"left" bson:"left"`
	Accountlink       string    `json:"accountlink" bson:"accountlink"`
	Things            string    `json:"things" bson:"things"`
	Members           string    `json:"members" bson:"members"`
	Unsend            string    `json:"unsend" bson:"unsend"`
	Vedioplaycomplete string    `json:"vedioplaycomplete" bson:"vedioplaycomplete"`
}

type Source struct {
	Type    string `json:"type" bson:"type"`
	UserId  string `json:"userid" bson:"userid"`
	RroupId string `json:"groupid" bson:"groupid"`
	RoomId  string `json:"roomid" bson:"roomid"`
}

type Emojis struct {
	Index     int    `json:"index" bson:"index"`
	Length    int    `json:"length" bson:"length"`
	ProductId string `json:"productid" bson:"productid"`
	EmojiId   string `json:"emojiid" bson:"emojiid"`
}

type Message struct {
	Id     string   `json:"id" bson:"id"`
	Text   string   `json:"text" bson:"text"`
	Emojis []Emojis `json:"emojis" bson:"emojis"`
	Metion string   `json:"metion" bson:"metion"`
}

// Pushmessage structure
type Pushmessage struct {
	Message string `mapstructure:"message"`
}
