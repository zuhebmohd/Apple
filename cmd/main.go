package main

import (
	"go-cluster-portal/db"
	"go-cluster-portal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
