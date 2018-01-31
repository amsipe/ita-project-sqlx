package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
func getAllProducts(c *gin.Context){
	products, err := fetchAllProducts()
	if err != nil {
		c.String(404,"Something went wrong")
		return
	}
	c.JSON(http.StatusOK,products)
}

func getAllOrders(c *gin.Context){
	orders, err := fetchAllOrders()
	if err != nil {
		c.String(404,"Something went wrong")
		return
	}
	c.JSON(http.StatusOK, orders)
}

func getOrderDetails(c *gin.Context){
	id := c.Param("id")
	details, err := fetchOneByOrderID(id)
	if err != nil {
		c.String(http.StatusNotFound, "OrderID not found")
		return
	}
	c.JSON(http.StatusOK, details)
}

func deleteOrder(c *gin.Context){
	id := c.Param("id")
	count, err := deleteOneByOrderID(id)
	if err != nil {
		c.String(http.StatusNotFound, "OrderID not found")
	}
	c.String(http.StatusOK, fmt.Sprintf("Rows removed: %v",count))
}

func addProduct(c *gin.Context){
	var product Product
	err := c.BindJSON(&product) 
	if err != nil {
		c.String(http.StatusNotFound, "Something went wrong.")
	}
	err = addOneProduct(&product)
	if err != nil {
		c.String(http.StatusOK,"product added")
	}

}

func addOrder(c *gin.Context){
	var order Order
}