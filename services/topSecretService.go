package services

import (
	"goCore/models"
)

type TopSecretService struct{}
var locationService LocationService
var messageService MessageService

func(l TopSecretService) ConstructResponse(topSecret models.TopSecret) (response models.TopSecretResponse){
	distancesArray := locationService.GetFloatArrayFromTopSecret(topSecret)
	x, y := locationService.GetLocation(distancesArray[0], distancesArray[1], distancesArray[2])
	secret := messageService.GetMessageFromTopSecret(topSecret)

	position := models.Position{
		X: x,
		Y: y,
	}
	secretResponse := models.TopSecretResponse{Position: position, Message: secret}

	return secretResponse
}