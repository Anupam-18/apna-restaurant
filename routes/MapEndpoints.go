package routes

import (
	"apna-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)
}

func MenuRoutes(router *gin.Engine) {
	router.POST("/menu/new", controllers.AddMenu)
	router.GET("/menu/all", controllers.GetAllMenus)
	router.GET("/menu/:id", controllers.GetMenuByID)
	router.PUT("/menu/update", controllers.UpdateMenu)
}
