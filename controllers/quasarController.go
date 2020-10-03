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

