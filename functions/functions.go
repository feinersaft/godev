package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "º K is ", celsius, "º C")
  
	celsius = kelvinToCelsius(233)
	fmt.Println("233 º K is ", celsius, "º C")
  
	celsius = celsiusToFahrenheit(26)
	fmt.Println("26 º C is ", celsius, "º F")
  
	celsius = kelvinToFahrenheit(233)
	fmt.Print("233 º K is ", celsius, "º F")
}
