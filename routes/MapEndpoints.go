package routes

import (
	"apna-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)
}
