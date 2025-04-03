package routes

import (
	"go-cluster-portal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/clusters", controllers.GetClusters)
	router.PUT("/clusters/:id", controllers.UpdateCluster)
	router.POST("/login", controllers.Login)
}
