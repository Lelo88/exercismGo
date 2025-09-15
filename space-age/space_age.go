package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var earthYear float64 = 31557600
	var planetYear float64
	switch planet {
	case "Earth":
		planetYear = 1
	case "Mercury":
		planetYear = 0.2408467
	case "Venus":
		planetYear = 0.61519726
	case "Mars":
		planetYear = 1.8808158
	case "Jupiter":
		planetYear = 11.862615
	case "Saturn":
		planetYear = 29.447498
	case "Uranus":
		planetYear = 84.016846
	case "Neptune":
		planetYear = 164.79132
	default:
		return -1
	}
	return seconds / (earthYear * planetYear)
}
