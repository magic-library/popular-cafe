package main

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine, config *AppConfig) {
	rc := NewRouterConfig(config)
	r.GET("/health-check", wrapper(healthCheck))
	r.GET("/coffee", wrapper(rc.CoffeeController.GetCoffeeList))
	r.POST("/take-coffee", wrapper(rc.CoffeeController.HandOutCoffee))
}

func wrapper(f func(c *gin.Context) (interface{}, retCode)) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, errCode := f(c)
		if errCode != Ok {
			c.JSON(statusCodes[errCode].Code, gin.H{"msg": statusCodes[errCode].Msg})
			return
		}
		c.JSON(statusCodes[Ok].Code, data)
	}
}
