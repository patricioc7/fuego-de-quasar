package services

import (
	"quasarFire/models"
	"testing"
)

func TestGetMessageFromTopSecret(t *testing.T) {
	service := MessageService{}
	kenobiSatellite := models.Satellite{Name: "kenobi", Distance: 989.945, Message: []string{"este", "", "", "mensaje", ""}}
	skywalkerSatellite := models.Satellite{Name: "skywalker", Distance: 608.276, Message: []string{"", "es", "", "", "secreto"}}
	satoSatellite := models.Satellite{Name: "sato", Distance: 500, Message: []string{"este", "", "un", "", ""}}
	secret := models.TopSecret{Satellites: [3]models.Satellite{kenobiSatellite, skywalkerSatellite, satoSatellite}}

	topSecretMessage := service.GetMessageFromTopSecret(secret)

	if topSecretMessage != "este es un mensaje secreto" {
		t.Errorf("Failed Interpolating message: " + topSecretMessage)
	}
}