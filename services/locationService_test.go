package services

import (
	"quasarFire/models"
	"testing"
)

func TestGetFloatArrayFromTopSecret(t *testing.T) {
	service := LocationService{}
	kenobiSatellite := models.Satellite{Name: "kenobi", Distance: 989.945, Message: []string{"este", "", "", "mensaje", ""}}
	skywalkerSatellite := models.Satellite{Name: "skywalker", Distance: 608.276, Message: []string{"", "es", "", "", "secreto"}}
	satoSatellite := models.Satellite{Name: "sato", Distance: 500, Message: []string{"este", "", "un", "", ""}}
	secret := models.TopSecret{Satellites: [3]models.Satellite{kenobiSatellite, skywalkerSatellite, satoSatellite}}

	x, y := service.GetFloatArrayFromTopSecret(secret)

	if x == nil && y == nil  {
		t.Errorf("Failed getting coordinates")
	}
}

func TestGetLocation(t *testing.T) {
	service := LocationService{}

	x, y := service.GetLocation(989.945, 608.276, 500 )

	if x == nil && y == nil  {

		t.Errorf("Failed getting coordinates")
	}
}

