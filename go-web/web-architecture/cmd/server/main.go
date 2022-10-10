package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/cmd/server/handler"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/pkg/store"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/internal/products"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db := store.NewStore(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Save())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
