package server

import (
	"github.com/gin-gonic/gin"
	"quasarFire/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("")
	{

		quasarController := new(controllers.QuasarController)

		v1.POST("/topsecret/", quasarController.FullPost)
		v1.POST("/topsecret_split/:satellite_name", quasarController.SplitPost)
		v1.GET("/topsecret_split/", quasarController.SplitGet)

	}

	return router
}