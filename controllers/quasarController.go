package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"quasarFire/models"
	"quasarFire/services"
)

type QuasarController struct{}

var locationService services.LocationService
var messageService services.MessageService
var topSecretService services.TopSecretService

func (u QuasarController) FullPost(c *gin.Context) {
	topSecret := models.TopSecret{}
	c.BindJSON(&topSecret)

	response, error := topSecretService.ConstructResponse(topSecret)
	fmt.Print(error)
	if error == nil {
		c.JSON(200, response)
	}else{
		c.JSON(400, gin.H{"error": error.Error()})
	}

	return
}

func (u QuasarController) SplitPost(c *gin.Context) {
	topSecret := models.TopSecretSplit{}
	c.BindJSON(&topSecret)
	name := c.Param("satellite_name")

	saved := topSecretService.SaveTopSecretSplitToCache(&topSecret, name)


	if saved {
		c.JSON(201, "Saved")

	}else{
		c.JSON(400, gin.H{"error": "bad request"})
	}
	return
}

func (u QuasarController) SplitGet(c *gin.Context) {
	topSecret := models.TopSecret{}
	c.BindJSON(&topSecret)

	response, error := topSecretService.ConstructResponseFromSplit()

	if error == nil {
		c.JSON(200, response)
	}else{
		c.JSON(428, gin.H{"error": error.Error()})
	}
	return
}

