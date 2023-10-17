package app

import (
	"apna-restaurant/middleware"
	"apna-restaurant/routes"
)

func mapUrls() {
	routes.AuthRoutes(router)
	router.Use(middleware.Authenticate())
	routes.MenuRoutes(router)
}
