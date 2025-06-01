package main

import (
	"asktils/db"
	"asktils/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.InitDB()

	routes.SetupRoutes(router, db)

	router.Run("127.0.0.1:8080")
}
