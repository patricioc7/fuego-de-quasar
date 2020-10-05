package services

import (
	"errors"
	"math"
	"quasarFire/models"
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

	const EPSILON = 0.01

	//kenobi
	const xPositionKenobi float64 = -500
	const yPositionKenobi float64 = -200

	//skywalker
	const xPositionSkywalker float64 = 100
	const yPositionSkywalker float64 = -100

	//sato
	const xPositionSato float64 = 500
	const yPositionSato float64 = 100

	var intersectionPointsLineAndCentersLineDisntance, xOffset, yOffset float64
	var xIntersectionPointsLineAndCentersLine, yIntersectionPointsLineAndCentersLine float64

	xDifferenceKenobiSkywalker := xPositionSkywalker - xPositionKenobi
	yDifferenceKenobiSkywalker := yPositionSkywalker - yPositionKenobi
	hKenobiSkywalker := math.Sqrt((yDifferenceKenobiSkywalker * yDifferenceKenobiSkywalker) + (xDifferenceKenobiSkywalker * xDifferenceKenobiSkywalker))
	if hKenobiSkywalker > (distanceKenobi + distanceSkywalker){
		//no solution
		return nil, nil
	}

	if hKenobiSkywalker < math.Abs(distanceKenobi-distanceSkywalker){
		// no solution. One circle contained in another
		return nil, nil
	}

	/* intersectionPointsLineAndCentersLine is the point where the line through the circle
	 * intersection points crosses the line between the circle
	 * centers.
	 */

	/* Determine the distance from Kenobi to intersectionPointsLineAndCentersLine. */
	intersectionPointsLineAndCentersLine := ((distanceKenobi * distanceKenobi) - (distanceSkywalker * distanceSkywalker) + (hKenobiSkywalker * hKenobiSkywalker)) / (2.0 * hKenobiSkywalker)

	/* Determine the coordinates of point 2. */
	xIntersectionPointsLineAndCentersLine = xPositionKenobi + (xDifferenceKenobiSkywalker * intersectionPointsLineAndCentersLine / hKenobiSkywalker)
	yIntersectionPointsLineAndCentersLine = yPositionKenobi + (yDifferenceKenobiSkywalker * intersectionPointsLineAndCentersLine / hKenobiSkywalker)

	/* Determine the distance from intersectionPointsLineAndCentersLine to either of the
	 * intersection points.
	 */
	intersectionPointsLineAndCentersLineDisntance = math.Sqrt((distanceKenobi * distanceKenobi) - (intersectionPointsLineAndCentersLine * intersectionPointsLineAndCentersLine))

	/* Determine the offsets of the intersection points from
	 * intersectionPointsLineAndCentersLine.
	 */
	xOffset = -yDifferenceKenobiSkywalker * (intersectionPointsLineAndCentersLineDisntance / hKenobiSkywalker)
	yOffset = xDifferenceKenobiSkywalker * (intersectionPointsLineAndCentersLineDisntance / hKenobiSkywalker)

	// Determine the absolute intersection points.
	intersectionPoint1X := xIntersectionPointsLineAndCentersLine + xOffset
	intersectionPoint2X := xIntersectionPointsLineAndCentersLine - xOffset
	intersectionPoint1Y := yIntersectionPointsLineAndCentersLine + yOffset
	intersectionPoint2Y := yIntersectionPointsLineAndCentersLine - yOffset

	//etermine if Sato's circle intersects at either of the above intersection points.
	xDifferenceKenobiSkywalker = intersectionPoint1X - xPositionSato
	yDifferenceKenobiSkywalker = intersectionPoint1Y - yPositionSato
	d1 := math.Sqrt((yDifferenceKenobiSkywalker * yDifferenceKenobiSkywalker) + (xDifferenceKenobiSkywalker * xDifferenceKenobiSkywalker))

	xDifferenceKenobiSkywalker = intersectionPoint2X - xPositionSato
	yDifferenceKenobiSkywalker = intersectionPoint2Y - yPositionSato
	d2 := math.Sqrt((yDifferenceKenobiSkywalker * yDifferenceKenobiSkywalker) + (xDifferenceKenobiSkywalker * xDifferenceKenobiSkywalker))


	if math.Abs(d1 - distanceSato) < EPSILON {
		x := float32(math.Round(intersectionPoint1X*100) / 100)
		y := float32(math.Round(intersectionPoint1Y*100) / 100)
		return &x, &y
	}
	if math.Abs(d2 - distanceSato) < EPSILON {

		x := float32(math.Round(intersectionPoint2X*100) / 100)
		y := float32(math.Round(intersectionPoint2Y*100) / 100)
		return &x, &y

	}
	return nil, nil
}
