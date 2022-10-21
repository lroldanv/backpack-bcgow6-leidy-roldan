package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	docs "github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/cmd/server/docs"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/cmd/server/handler"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go-web/web-architecture/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/products"
)

// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @description     This is a simple API to handle products
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

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
	api := r.Group("api/v1")

	// Swagger docs
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middlewares
	api.Use(handler.TokenMiddleware())

	pr := api.Group("/products")
	{
		pr.POST("/", p.Save())
		pr.GET("/", p.GetAll())
		pr.PUT("/:id", p.Update())
		pr.DELETE("/:id", p.Delete())

	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
