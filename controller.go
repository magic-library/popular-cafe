package main

import "github.com/gin-gonic/gin"

func healthCheck(c *gin.Context) (interface{}, retCode) {
	return "ok", Ok
}

type CoffeeController struct {
	Table *Table
}

func (controller *CoffeeController) GetCoffeeList(c *gin.Context) (interface{}, retCode) {
	res := CoffeeResponse{
		Coffee: []*Coffee{},
	}
	return res, Ok
}

func (controller *CoffeeController) HandOutCoffee(c *gin.Context) (interface{}, retCode) {
	return "give it to you", Ok
}
