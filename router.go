package main

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/products", getAllProducts)
	r.GET("/orders", getAllOrders)
	r.GET("/orders/:id", getOrderDetails)
	r.DELETE("/orders/:id", deleteOrder)
	r.POST("/products", addProduct)
	return r
}