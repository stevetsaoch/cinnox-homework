package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/stevetsaoch/cinnox-homework/config"
	"github.com/stevetsaoch/cinnox-homework/model"
	"go.mongodb.org/mongo-driver/bson"
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

func PushMessage(c *gin.Context) {
	// declare message
	var message model.Pushmessage
	err := c.Bind(&message)

	// load config
	config_, err := config.Loadconfig()

	if err != nil {
		log.Fatal(err)
	}

	// mongodb instance
	client_mongo := config.ConnectDB()
	bot := config.LineBot()

	// collection instance
	collection := client_mongo.Database(config_.DatabaseName).Collection(config_.CollectionName)

	// get all distinct userid for sending promote message
	results, err := collection.Distinct(context.TODO(), "source.userid", bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		if val, ok := result.(string); ok {
			if _, err := bot.PushMessage(val, linebot.NewTextMessage(message.Message)).Do(); err != nil {
				fmt.Printf("Message not sent!")
				log.Fatal(err)
			}
		}
	}

	// return userid of receiving message
	c.JSON(http.StatusOK, results)
}

func GetAllMessages(c *gin.Context) {
	var linemessage model.LineEvent
	var results []model.LineEvent
	// load config
	config_, err := config.Loadconfig()

	if err != nil {
		log.Fatal(err)
	}

	// database and collection instance
	client := config.ConnectDB()
	coll := client.Database(config_.DatabaseName).Collection(config_.CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// set a cursor
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// loop through all results and save in results slice
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&linemessage); err != nil {
			log.Fatal(err)
		}
		results = append(results, linemessage)
	}

	c.JSON(http.StatusOK, results)
}
