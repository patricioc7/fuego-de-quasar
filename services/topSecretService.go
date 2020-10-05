package services

import (
	"errors"
	"quasarFire/cache"
	"quasarFire/models"
)

type TopSecretService struct{}
var locationService LocationService
var messageService MessageService


func(l TopSecretService) ConstructResponse(topSecret models.TopSecret) (*models.TopSecretResponse, error){

	distancesArray, error := locationService.GetFloatArrayFromTopSecret(topSecret)
	if error == nil {

		//Using GoRoutines here is unnecessary, but is fun to use it:
		positionChannel := make(chan *models.Position,  2)
		secretMessageChannel := make(chan string, 1)
		go func() {
			x, y := locationService.GetLocation(distancesArray[0], distancesArray[1], distancesArray[2])
			positionChannel <- &models.Position{X: *x, Y: *y}
		}()
		go func() {
			secret := messageService.GetMessageFromTopSecret(topSecret)
			secretMessageChannel <- secret
		}()

		position := <-positionChannel
		secret := <- secretMessageChannel
		if secret != "" && position != nil {
			secretResponse := models.TopSecretResponse{Position: *position, Message: secret}
			return &secretResponse, nil
		}else{
			error = errors.New("could not calculate position or message with provided info")
		}

	}
	if error == nil {
		error = errors.New("could not calculate position or message with provided info")
	}
	return nil, error

}

func(l TopSecretService) SaveTopSecretSplitToCache(topSecretSplit *models.TopSecretSplit, name string) (saved bool)  {
	if topSecretSplit != nil {
		cache.SetValue(name, topSecretSplit)
		return true
	}
	return false

}

func(l TopSecretService) ConstructResponseFromSplit() (**models.TopSecretResponse, error) {

	kenobiRequest, kenobiFound := cache.GetValue("kenobi")
	skywalkerRequest, skywalkerFound := cache.GetValue("skywalker")
	satoRequest, satoFound := cache.GetValue("sato")

	if kenobiFound  && skywalkerFound && satoFound{

		parsedKenobi := kenobiRequest.(*models.TopSecretSplit)
		parsedSkywalker := skywalkerRequest.(*models.TopSecretSplit)
		parsedSato := satoRequest.(*models.TopSecretSplit)

		kenobiSatellite := models.Satellite{Name: "kenobi", Distance: parsedKenobi.Distance, Message: parsedKenobi.Message}
		skywalkerSatellite := models.Satellite{Name: "skywalker", Distance: parsedSkywalker.Distance, Message: parsedSkywalker.Message}
		satoSatellite := models.Satellite{Name: "sato", Distance: parsedSato.Distance, Message: parsedSato.Message}

		satellites := [3]models.Satellite{kenobiSatellite, skywalkerSatellite, satoSatellite}

		topSecretReadyToProcess := models.TopSecret{Satellites: satellites}

		response, error := l.ConstructResponse(topSecretReadyToProcess)
		if error != nil {
			return nil, error
		}
		return &response, nil

	}
	return nil, errors.New("not enough info yet")
}