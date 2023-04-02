package api

import (
	"example/line-bot-ledger/utils"

	"github.com/gin-gonic/gin"
)

func Router(NODE_ENV string) *gin.Engine {
	if NODE_ENV != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/healthcheck", healthcheck)

	router.Use(utils.JwtMiddleware())
	return router
}
