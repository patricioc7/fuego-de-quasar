package server

import (
	"github.com/gin-gonic/gin"
	"goCore/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//adding jwt
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		quasarController := new(controllers.QuasarController)

		v1.POST("/secretmessage", quasarController.Test)

	}

	return router
}