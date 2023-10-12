package app

import (
	"apna-restaurant/database"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func init() {
	database.DBConfig()
}
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
