package services

import (
	"quasarFire/cache"
	"quasarFire/models"
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

func(l TopSecretService) SaveTopSecretSplitToCache(topSecret models.TopSecretSplit, name string){

	cache.SetValue(name, topSecret)
}

func(l TopSecretService) ConstructResponseFromSplit() *models.TopSecretResponse {

	kenobiRequest, kenobiFound := cache.GetValue("kenobi")
	skywalkerRequest, skywalkerFound := cache.GetValue("skywalker")
	satoRequest, satoFound := cache.GetValue("sato")

	if kenobiFound  && skywalkerFound && satoFound{

		parsedKenobi := kenobiRequest.(models.TopSecretSplit)
		parsedSkywalker := skywalkerRequest.(models.TopSecretSplit)
		parsedSato := satoRequest.(models.TopSecretSplit)

		kenobiSatellite := models.Satellite{Name: "kenobi", Distance: parsedKenobi.Distance, Message: parsedKenobi.Message}
		skywalkerSatellite := models.Satellite{Name: "skywalker", Distance: parsedSkywalker.Distance, Message: parsedSkywalker.Message}
		satoSatellite := models.Satellite{Name: "sato", Distance: parsedSato.Distance, Message: parsedSato.Message}

		satellites := [3]models.Satellite{kenobiSatellite, skywalkerSatellite, satoSatellite}

		topSecretReadyToProcess := models.TopSecret{Satellites: satellites}

		response := l.ConstructResponse(topSecretReadyToProcess)
		return &response

	}
	return nil
}