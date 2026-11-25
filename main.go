package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ujjwalsh-2209/event-booking/db"
	"github.com/ujjwalsh-2209/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8081") //localhost:8081

}
