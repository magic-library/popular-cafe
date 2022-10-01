package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	appConfig := NewAppConfig()
	Route(router, appConfig)
	router.Run(":8888")
}
