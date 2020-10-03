package controllers

import (
	"github.com/gin-gonic/gin"
	"goCore/models"
	"goCore/services"
)

type QuasarController struct{}

var locationService services.LocationService
var messageService services.MessageService
var topSecretService services.TopSecretService

func (u QuasarController) Test(c *gin.Context) {
	topSecret := models.TopSecret{}
	c.BindJSON(&topSecret)

	response := topSecretService.ConstructResponse(topSecret)

	//handle errors

	c.JSON(200, response)
	return
}

func (u QuasarController) SplitPost(c *gin.Context) {
	topSecret := models.TopSecretSplit{}
	c.BindJSON(&topSecret)
	name := c.Param("satellite_name")

	topSecretService.SaveTopSecretSplitToCache(topSecret, name)

	//handle errors

	c.JSON(201, "Saved")
	return
}

func (u QuasarController) SplitGet(c *gin.Context) {
	topSecret := models.TopSecret{}
	c.BindJSON(&topSecret)

	response := topSecretService.ConstructResponseFromSplit()

	if response != nil {
		c.JSON(200, response)
	}else{
		c.JSON(428, "Not enough info yet")
	}
	return
}

