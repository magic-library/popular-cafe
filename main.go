package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magic-library/popular-cafe/api"
)

func main() {
	router := gin.Default()

	api.Route(router)
	router.Run(":8888")
}
