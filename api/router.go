package api

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.GET("/health-check", wrapper(healthCheck))
	r.GET("/coffee", wrapper(getCoffeeList))
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
