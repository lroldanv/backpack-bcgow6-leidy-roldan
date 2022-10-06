package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID        int       // `json:"id"`
	Name      string    // `json:"name"`
	Color     string    // `json:"color"`
	Price     float64   // `json:"price"`
	Stock     uint      // `json:"stock"`
	Code      string    //`json:"code"`
	Published bool      // `json:"published"`
	CreatedAt time.Time // `json:"createdAt"`
}

func (p Product) GetProductsFromFile(filePath string) (*[]Product, error) {
	var products []Product

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return &products, err
	}

	err = json.Unmarshal(content, &products)
	if err != nil {
		log.Fatal(err)
		return &products, err
	}

	return &products, nil
}

func main() {

	p := Product{}

	products, err := p.GetProductsFromFile("./go-web/products.json")
	fmt.Println(products)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create a router with gin
	router := gin.Default()

	// GET: localhost:8080/hello
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Joha",
		})
	})

	// router.GET("/movies/:id", func(c *gin.Context) {
	//     name := c.Param("id")
	//     c.String(http.StatusOK, "Movie %s", id)
	// })
	// router.Run(":8080")

	// Group routes
	product := router.Group("/products")
	{
		// GET: localhost:8080/products
		product.GET("/", GetAll(products))
		// product.GET("/search", SearchProduct)
		// GET: localhost:8080/products/1
		product.GET("/:id", GetProductByID(products))
	}

	router.Run() // Runs in 8080 by default

}

func GetAll(products *[]Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, products) // use IndentedJSON only for development purposes
	}
}

func GetProductByID(products *[]Product) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, product := range *products {
			if product.ID == id {
				c.IndentedJSON(http.StatusOK, product) // use IndentedJSON only for development purposes
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
	}
}

func GetProductsByColor(color string, products []Product) []*Product {
	var filtered []*Product
	for _, product := range products {
		if product.Color == color {
			filtered = append(filtered, &product)
		}
	}
	return filtered
}
