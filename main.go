package main

import (
	"go-lang-restapi/db"
	"go-lang-restapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	defer db.DB.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
