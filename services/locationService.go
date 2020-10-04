package services

import (
	"fmt"
	"quasarFire/models"
	"math"
)
type LocationService struct{}

func(l LocationService) GetFloatArrayFromTopSecret(topSecret models.TopSecret) (distances [3]float32){
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

	return result

}


func (l LocationService) GetLocation(distances ...float32) (x, y float32){



	r0 := float64(distances[0])
	r1:= float64(distances[1])
	r2 := float64(distances[2])

	EPSILON := 0.01

	//kenobi
	var x0  float64 = -500
	var y0 float64 = -200

	//skywalker
	var x1 float64 = 100
	var y1 float64 = -100

	//sato
	var x2 float64 = 500
	var y2 float64 = 100

	var a, dx, dy, d, h, rx, ry float64
	var point2X, point2Y float64

	dx = x1 - x0
	dy = y1 - y0
	d = math.Sqrt((dy * dy) + (dx * dx))
	if d > (r0 + r1){
		/* no solution. circles do not intersect. */
		return 0, 0
	}

	if d < math.Abs(r0 - r1){
		/* no solution. one circle is contained in the other */
		return 0, 0
	}

	/* 'point 2' is the point where the line through the circle
	 * intersection points crosses the line between the circle
	 * centers.
	 */

	/* Determine the distance from point 0 to point 2. */

	a = ((r0*r0) - (r1*r1) + (d*d)) / (2.0 * d)

	/* Determine the coordinates of point 2. */
	point2X = x0 + (dx * a/d)
	point2Y = y0 + (dy * a/d)

	/* Determine the distance from point 2 to either of the
	 * intersection points.
	 */
	h = math.Sqrt((r0*r0) - (a*a))

	/* Now determine the offsets of the intersection points from
	 * point 2.
	 */
	rx = -dy * (h/d)
	ry = dx * (h/d)

	/* Determine the absolute intersection points. */
	intersectionPoint1X := point2X + rx
	intersectionPoint2X := point2X - rx
	intersectionPoint1Y := point2Y + ry
	intersectionPoint2Y := point2Y - ry

	/* Lets determine if circle 3 intersects at either of the above intersection points. */
	dx = intersectionPoint1X - x2
	dy = intersectionPoint1Y - y2
	d1 := math.Sqrt((dy*dy) + (dx*dx))

	dx = intersectionPoint2X - x2
	dy = intersectionPoint2Y - y2
	d2 := math.Sqrt((dy*dy) + (dx*dx))

	fmt.Println(math.Abs(d1 - r2))
	fmt.Println(math.Abs(d2 - r2))

	if math.Abs(d1 - r2) < EPSILON {
		fmt.Println("1 INTERSECTION Circle1 AND Circle2 AND Circle3:" + "(" + fmt.Sprintf("%f", math.Round(intersectionPoint1X*100)/100 )  + "," + fmt.Sprintf("%f", math.Round(intersectionPoint1Y*100)/100 )  + ")")
		return float32(math.Round(intersectionPoint1X*100)/100), float32(math.Round(intersectionPoint1Y*100)/100) ;
	} else if math.Abs(d2 - r2) < EPSILON {
		fmt.Println("2 INTERSECTION Circle1 AND Circle2 AND Circle3:" + "(" + fmt.Sprintf("%f", intersectionPoint2X)  + "," + fmt.Sprintf("%f", intersectionPoint2Y) + ")") //here was an error

	} else {
		fmt.Println("INTERSECTION Circle1 AND Circle2 AND Circle3:" + "NONE")
	}
	return 0, 0;
}
