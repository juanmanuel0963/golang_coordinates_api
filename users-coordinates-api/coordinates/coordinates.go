package coordinates

import (
	"math"
	"sfox/v1/users-coordinates-api/config"
	"sfox/v1/users-coordinates-api/models"
)

// Convert degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// Calculate the distance between two coordinates using the Haversine formula
func HaversineDistance(coord1, coord2 models.UserCoordinates, config *config.Configuration) float64 {

	//The values for Latitude and Longitude are converted from degrees to radians using the DegreesToRadians function to perform trigonometric calculations.
	lat1 := DegreesToRadians(coord1.Latitude)
	lon1 := DegreesToRadians(coord1.Longitude)
	lat2 := DegreesToRadians(coord2.Latitude)
	lon2 := DegreesToRadians(coord2.Longitude)

	//dLat is the difference in latitude between the two sets of coordinates in radians, and dLon is the difference in longitude in radians.
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	//The Haversine formula is applied to calculate the great-circle distance between the two coordinates. It consists of several steps:
	//Calculate the square of half the difference in latitudes (math.Pow(math.Sin(dLat/2), 2)).
	//Calculate the square of half the difference in longitudes (math.Pow(math.Sin(dLon/2), 2)).
	//Calculate the product of the cosine of lat1 and the cosine of lat2 (math.Cos(lat1)*math.Cos(lat2)).
	//Calculate a by summing the results of the three previous calculations.
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)

	//Calculate c as 2 times the arctangent of the square root of a divided by the square root of 1 - a.
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	//Returns the calculated distance, which is the Earth's radius multiplied by c. The result is in kilometers (or the unit of measurement used for earthRadius).
	return config.EarthRadius * c
}

// Filter coordinates by distance to a reference point
func FilterCoordinatesByDistance(referencePoint models.UserCoordinates, inputUserCoordinates []models.UserCoordinates, maxDistance float64, config *config.Configuration) []models.UserCoordinates {

	var outputUserCoordinates []models.UserCoordinates

	//Iterate the Users Coordinates list
	for _, userCoordinates := range inputUserCoordinates {

		// Calculate the distance between two coordinates using the Haversine formula
		distance := HaversineDistance(referencePoint, userCoordinates, config)

		//If the distance is minor than the max expected distance
		if distance <= maxDistance {

			userCoordinates.Distance = distance

			//Add the point to the output list
			outputUserCoordinates = append(outputUserCoordinates, userCoordinates)
			//fmt.Printf("Id: %v, User: %s, Distance: %v km\n", userCoordinates.Id, userCoordinates.Name, distance)
		}

	}

	return outputUserCoordinates

}
