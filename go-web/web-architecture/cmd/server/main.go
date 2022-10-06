package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/cmd/server/handler"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Save())
	pr.GET("/", p.GetAll())
	pr.PUT("/id", p.Update())

	r.Run()
}
