package routers

import (
	"VincentLimarus/log-activity/controllers/services"
	"VincentLimarus/log-activity/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutersConfiguration() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"Code": 400, "message": "Route not found"})
	})

	base := router.Group("api/v1")
	services.BaseUserService(base)

	auth := router.Group("api/v1")
	auth.Use(middlewares.RequiredAuth())

	services.AuthOrderService(auth)	

	return router

}