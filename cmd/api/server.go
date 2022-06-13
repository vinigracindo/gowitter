package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(port int) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return router
}
