package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevetsaoch/cinnox-homework/controller"
)

func LineRoutes(router *gin.Engine) {
	router.POST("/callback", controller.ReceiveMessage)
	router.POST("/push", controller.PushMessage)
	router.GET("/all", controller.GetAllMessages)
}
