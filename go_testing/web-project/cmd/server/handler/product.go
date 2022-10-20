package handler

import (
	"fmt"

	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/products"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/web"
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

// Products godoc
// @Summary  Show the list of products
// @Tags     Products
// @Produce  json
// @Param    token  header    string        true  "Token"
// @Success  200    {object}  web.Response  "List of products"
// @Failure  401    {object}  web.Response  "Unauthorized + invalid token"
// @Failure  500    {object}  web.Response  "Internal server error "
// @Failure  404    {object}  web.Response  "Not found"
// @Router   /products [GET]
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

// Update product godoc
// @Summary  Update a product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "Id product"
// @Param    token    header    string          true  "Token"
// @Param    product  body      Request  true   "Product to update"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  404      {object}  web.Response
// @Router   /products/{id} [PUT]
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

		// TODO middleware to assign a creation date

		product, err := p.service.Update(int(id), req.Name, req.Color, req.Code, req.Price, req.Stock, req.Published, req.CreatedAt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}

// Update a product name godoc
// @Summary      Update a product name
// @Tags         Products
// @Accept       json
// @Produce      json
// @Description  This endpoint update the product name
// @Param        token  header    string            true  "Token"
// @Param        id     path      int               true  "Product Id"
// @Param        name   body      NameRequest		true  "Product name"
// @Success      200    {object}  web.Response
// @Failure      401    {object}  web.Response
// @Failure      400    {object}  web.Response
// @Failure      404    {object}  web.Response
// @Router       /products/{id} [PATCH]
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

// Save product godoc
// @Summary  Store a product in db
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "Token"
// @Param    product  body      Request  true  "Product to Store"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Router   /products [POST]
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

// Delete Product
// @Summary  Delete product
// @Tags     Products
// @Param    id     path      int     true  "Product id"
// @Param    token  header    string  true  "Token"
// @Success  204
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /products/{id} [DELETE]
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
