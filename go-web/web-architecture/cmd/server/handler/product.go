package handler

import (
	"fmt"

	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/internal/products"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/pkg/web"
)

type Request struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Color     string    `json:"color" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Stock     uint      `json:"stock" binding:"required"`
	Code      string    `json:"code" binding:"required"`
	Published bool      `json:"published" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type NameRequest struct {
	Name string `json:"name" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, products, ""))
	}
}

func (p *Product) GetProductByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		for _, product := range products {
			if product.ID == id {
				ctx.JSON(200, web.NewResponse(200, product, ""))
				return
			}
		}
		ctx.JSON(400, web.NewResponse(404, nil, fmt.Sprintf("Product with id: %d does not exist", id)))
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid ID"))
			return
		}

		var req Request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product color is required"))
			return
		}
		if req.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product code is required"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}

		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}

		// TODO validate creation date

		product, err := p.service.Update(int(id), req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published, req.CreatedAt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID"+err.Error()))
			return
		}

		var req NameRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		product, err := p.service.UpdateName(id, req.Name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, product, ""))

	}
}

func (p *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req Request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product color is required"))
			return
		}
		if req.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Product code is required"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}

		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Product name is required"))
			return
		}

		product, err := p.service.Save(req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID"+err.Error()))
			return
		}

		if err := p.service.Delete(id); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nil, ""))
	}
}
