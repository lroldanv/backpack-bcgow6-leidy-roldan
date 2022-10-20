package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/web"
)

func TokenMiddleware() gin.HandlerFunc {
	tokenAPI := os.Getenv("TOKEN")
	if tokenAPI == "" {
		log.Fatal("A token has not been registered")
	}

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if tokenAPI != token {
			ctx.AbortWithStatusJSON(403, web.NewResponse(403, nil, "Forbbiden"))
			return
		}
		ctx.Next()
	}
}
