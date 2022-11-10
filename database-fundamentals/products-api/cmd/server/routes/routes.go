package routes

import (
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/cmd/server/handler"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/product"

	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildProductRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)
	//r.rg.GET("/movies", handler.GetAll())
	r.rg.GET("/products", handler.GetByName())
	//r.rg.POST("/movies", handler.Create())
	//r.rg.DELETE("/movies/:id", handler.Delete())
	//r.rg.PATCH("/movies/:id", handler.Update())
}
