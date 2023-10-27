package main

import (
	// "fmt"
	// "log"
	"ami_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}

}

func getProduct(c *gin.Context) {
	code := c.Param("id")
	product := models.GetProduct(code)

	if product == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, product)

	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
		})
	})
	router.GET("/products", getProducts)
	router.GET("/product/:id", getProduct)
	router.POST("/products", addProduct)
	router.Run("localhost:7070")
}

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homeLink)
// 	log.Fatal(http.ListenAndServe(":7070", router))
// }
