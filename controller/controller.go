package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/stevetsaoch/cinnox-homework/config"
)

func ReceiveMessage(c *gin.Context) {
	// load config
	config_, err := config.Loadconfig()

	if err != nil {
		log.Fatal(err)
	}

	// connect to database and collection
	client_mongo := config.ConnectDB()
	collection := client_mongo.Database(config_.DatabaseName).Collection(config_.CollectionName)

	// declare context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// declare line bot instance
	bot := config.LineBot()

	// parse request
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
	}

	// reply message
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// send back message
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Message received")).Do(); err != nil {
				log.Println(err.Error())
			}
		}
		// save event to mongodb
		res, err := collection.InsertOne(ctx, event)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}

}
