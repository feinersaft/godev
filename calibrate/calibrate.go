package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}
func main() {
	var offset kelvin = 5.0
	sensor := calibrate(realSensor, offset)
	offset = 55.0 // Modifying the variable after passing it to the function does not change the value passed to the function.
	fmt.Println(sensor())
	for i := 0; i < 5; i++ {
		// Call the new sensor function multiple times and notice that the original
		// fakeSensor is still being called each time, resulting in random values.
		sensor = calibrate(fakeSensor, offset)
		fmt.Println(sensor())
	}

}
