package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stevetsaoch/cinnox-homework/routes"
)

func main() {
	// initialize server
	router := gin.Default()

	// add routes
	routes.LineRoutes(router)

	// run server
	router.Run("localhost:8000")
}
