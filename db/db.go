package db

import (
	"go-cluster-portal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "user=your_database_user dbname=clusters_db sslmode=disable")
	if err != nil {
		panic("failed to connect to database")
	}
	DB.AutoMigrate(&models.Cluster{}, &models.User{})
}
