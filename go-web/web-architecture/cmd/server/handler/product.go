package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/internal/products"
)

type request struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Color     string    `json:"color" validate:"required"`
	Price     float64   `json:"price" validate:"required"`
	Stock     uint      `json:"stock" validate:"required"`
	Code      string    `json:"code" validate:"required"`
	Published bool      `json:"published" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

		}
		ctx.IndentedJSON(http.StatusOK, products) // use IndentedJSON only for development purposes
	}
}

func (p *Product) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, product := range products {
			if product.ID == id {
				ctx.JSON(http.StatusOK, product)
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El campo nombre es requerido"})
			return
		}
		if req.Color == "" {
			ctx.JSON(400, gin.H{"error": "El campo color es requerido"})
			return
		}
		if req.Code == "" {
			ctx.JSON(400, gin.H{"error": "El campo code es requerido"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "El campo price es requerido"})
			return
		}

		if req.Stock == 0 {
			ctx.JSON(400, gin.H{"error": "El campo stock es requerido"})
			return
		}

		// if !req.Published {
		// 	ctx.JSON(400, gin.H{"error": "El campo stock es requerido"})
		// 	return
		// }

		// TODO validate creation date

		p, err := c.service.Update(int(id), req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published, req.CreatedAt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}

}
