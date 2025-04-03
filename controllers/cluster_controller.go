package controllers

import (
	"go-cluster-portal/db"
	"go-cluster-portal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClusters(c *gin.Context) {
	var clusters []models.Cluster
	db.DB.Find(&clusters)
	c.JSON(http.StatusOK, clusters)
}

func UpdateCluster(c *gin.Context) {
	var cluster models.Cluster
	if err := db.DB.Where("id = ?", c.Param("id")).First(&cluster).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cluster not found"})
		return
	}
	if err := c.ShouldBindJSON(&cluster); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	db.DB.Save(&cluster)
	c.JSON(http.StatusOK, cluster)
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	var foundUser models.User
	if err := db.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&foundUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "role": foundUser.Role})
}
