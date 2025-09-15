package meteorology

import "strconv"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

func (tu TemperatureUnit) String() string {
	switch tu {
	case Celsius:
		return string('°') + "C"
	case Fahrenheit:
		return string('°') + "F"
	}
	return ""
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

func (t Temperature) String() string {
	return strconv.Itoa(t.degree) + " " + t.unit.String()
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	}
	return ""
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (s Speed) String() string {
	return strconv.Itoa(s.magnitude)+ " " + s.unit.String()
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (md MeteorologyData) String() string {
    return md.location + ": " +
        md.temperature.String() + ", Wind " +
        md.windDirection + " at " + md.windSpeed.String() + ", " +
        strconv.Itoa(md.humidity) + "%" + " Humidity"
}