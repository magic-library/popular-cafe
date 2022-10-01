package api

import "github.com/gin-gonic/gin"

func healthCheck(c *gin.Context) (interface{}, retCode) {
	return "ok", Ok
}

func getCoffeeList(c *gin.Context) (interface{}, retCode) {
	res := CoffeeResponse{
		Coffee: []*Coffee{},
	}
	return res, Ok
}
