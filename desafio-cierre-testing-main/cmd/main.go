package main

import (
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/desafio-cierre-testing-main/cmd/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if err := r.Run(":18085"); err != nil {
		panic(err)
	}
}
