package main

import (
	"log"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/cmd/server/routes"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
