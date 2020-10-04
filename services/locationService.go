package services

import (
	"errors"
	"fmt"
	"quasarFire/models"
	"math"
)
type LocationService struct{}

func(l LocationService) GetFloatArrayFromTopSecret(topSecret models.TopSecret) (*[3]float32, error) {
	if len(topSecret.Satellites)<3  {
		return nil, errors.New("not enough satellites")
	}
	var result [3]float32
	for i:= 0; i < len(topSecret.Satellites); i+=1{
		switch topSecret.Satellites[i].Name {
		case "kenobi":
			result[0] = topSecret.Satellites[i].Distance
		case "skywalker":
			result[1] = topSecret.Satellites[i].Distance
		case "sato":
			result[2] = topSecret.Satellites[i].Distance
		}
	}

	return &result, nil

}


func (l LocationService) GetLocation(distances ...float32) (x, y *float32) {

	distanceKenobi := float64(distances[0])
	distanceSkywalker := float64(distances[1])
	distanceSato := float64(distances[2])

	EPSILON := 0.01

	//kenobi
	var xKenobi float64 = -500
	var yKenobi float64 = -200

	//skywalker
	var xSkywalker float64 = 100
	var ySkywalker float64 = -100

	//sato
	var xSato float64 = 500
	var ySato float64 = 100

	distanceXSkywalkerKenobi := xSkywalker - xKenobi
	distanceYSkywalkerKenobi := ySkywalker - yKenobi
	hipotSkywalkerKenobi := math.Sqrt((distanceYSkywalkerKenobi * distanceYSkywalkerKenobi) + (distanceXSkywalkerKenobi * distanceXSkywalkerKenobi))
	if hipotSkywalkerKenobi > (distanceKenobi + distanceSkywalker){
		//No solution
		return nil, nil
	}

	if hipotSkywalkerKenobi < math.Abs(distanceKenobi-distanceSkywalker){
		//No solution
		return nil, nil
	}

	distanceToFirstIntersectionPoint := ((distanceKenobi * distanceKenobi) - (distanceSkywalker * distanceSkywalker) +
		(hipotSkywalkerKenobi * hipotSkywalkerKenobi)) / (2.0 * hipotSkywalkerKenobi)

	intersectionPointX := xKenobi + (distanceXSkywalkerKenobi * distanceToFirstIntersectionPoint / hipotSkywalkerKenobi)
	intersectionPointY := yKenobi + (distanceYSkywalkerKenobi * distanceToFirstIntersectionPoint / hipotSkywalkerKenobi)


	dFromFirstInterPointToOthersInterPoints := math.Sqrt((distanceKenobi * distanceKenobi) - (distanceToFirstIntersectionPoint * distanceToFirstIntersectionPoint))

	intersectionPointCompensationX := -distanceYSkywalkerKenobi * (dFromFirstInterPointToOthersInterPoints / hipotSkywalkerKenobi)
	intersectionPointCompensationY := distanceXSkywalkerKenobi * (dFromFirstInterPointToOthersInterPoints / hipotSkywalkerKenobi)

	intersectionPoint1X := intersectionPointX + intersectionPointCompensationX
	intersectionPoint1Y := intersectionPointY + intersectionPointCompensationY

	distanceXSkywalkerKenobi = intersectionPoint1X - xSato
	distanceYSkywalkerKenobi = intersectionPoint1Y - ySato
	d1 := math.Sqrt((distanceYSkywalkerKenobi * distanceYSkywalkerKenobi) + (distanceXSkywalkerKenobi * distanceXSkywalkerKenobi))

	if math.Abs(d1 -distanceSato) < EPSILON {
		x := float32(math.Round(intersectionPoint1X*100) / 100)
		y := float32(math.Round(intersectionPoint1Y*100) / 100)
		fmt.Println("Intersection found" + "(" + fmt.Sprintf("%f", math.Round(intersectionPoint1X*100)/100 )  + "," + fmt.Sprintf("%f", math.Round(intersectionPoint1Y*100)/100 )  + ")")
		return &x, &y
	}
	return nil, nil
}
