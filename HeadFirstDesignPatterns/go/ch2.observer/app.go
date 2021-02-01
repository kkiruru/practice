package main

import (
	"fmt"

	"headfirst.design.patterns/observer/display"
	"headfirst.design.patterns/observer/weatherdata"
)

func main() {
	weatherData := weatherdata.WeatherData{}

	display.NewCurrentConditionsDisplay(&weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)

	fmt.Println("end")
}
