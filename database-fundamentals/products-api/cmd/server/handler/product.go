package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/product"
)

type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{
		service: service,
	}
}

func (p *Product) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		name := ctx.Query("name")

		product, err := p.service.GetByName(ctx, name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		ctx.JSON(200, gin.H{"product": product})
	}
}
