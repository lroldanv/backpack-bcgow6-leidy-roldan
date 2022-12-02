package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/product"
)

type requestName struct {
	Name string `json:"name" binding:"required"`
}
type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{
		service: service,
	}
}

// func (p *Product) GetByName() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		name := ctx.Param("name")

// 		product, err := p.service.GetByName(ctx, name)
// 		if err != nil {
// 			ctx.JSON(http.StatusNotFound, err)
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{"product": product})
// 	}
// }

func (p *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestName
		if err := c.ShouldBindJSON(&req); err != nil {
			// if strings.Contains(err.Error(), "'required' tag") {
			// 	c.JSON(http.StatusUnprocessableEntity, err.Error())
			// 	return
			// }

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := p.service.GetByName(c, req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"product": product})
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		product, err := p.service.Store(ctx, product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, gin.H{"product": product.Name + " added"})
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = p.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{"delete": id})
	}
}
