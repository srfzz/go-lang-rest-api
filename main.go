package main

import (
	"go-lang-restapi/db"
	"go-lang-restapi/routes"
	"go-lang-restapi/routes/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	defer db.DB.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)
	auth.RegisterAuthRoutes(server)
	server.Run(":8080")
}
