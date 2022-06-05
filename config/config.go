package config

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"
	"github.com/stevetsaoch/cinnox-homework/model"
)

func Loadconfig() (config model.Config, err error) {
	path := "/home/stevetsaoch/go_projects/cinnox-homework/config/"
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(path)
	vp.AutomaticEnv()

	err = vp.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	err = vp.Unmarshal(&config)

	return
}

// return a linebot instance
func LineBot() (client *linebot.Client) {
	// load config data
	config, err := Loadconfig()
	if err != nil {
		log.Fatal(err)
	}

	// create line client
	client, err = linebot.New(config.LineChannelSecret, config.LineChannelAccessToken)

	if err != nil {
		log.Fatal(err)
	}

	return
}

// return URI for login to mongodb
func MongoURI() (URI string) {

	// load config file
	config, err := Loadconfig()
	if err != nil {
		log.Fatal(err)
	}

	// build mogdb URI
	URI = "mongodb://" + config.MongoUser + ":" + config.MongoPwd + "@127.0.0.1:2717/?authSource=admin"
	return
}
