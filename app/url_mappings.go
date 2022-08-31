package app

import (
	"resttest/controllers/ping"
	controllers "resttest/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
}
