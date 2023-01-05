package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((9.0 / 5.0 * k) - 459.67)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((9.0 / 5.0 * c) + 32)
}
func (f fahrenheit) kelvin() kelvin {
	return kelvin(5.0 / 9.0 * (f + 459.67))
}
func (f fahrenheit) celsius() celsius {
	return celsius(5.0 / 9.0 * (f - 32))
}
func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Println(k, "º K is ", c, "º C")
	f := k.fahrenheit()
	fmt.Println(k, "° K is ", f, "° F")
	c = f.celsius()
	fmt.Println(f, "º F is ", c, "º C")
	k = f.kelvin()
	fmt.Println(f, "° F is ", k, "° K")
	f = c.fahrenheit()
	fmt.Println(c, "º C is ", f, "º F")
	k = c.kelvin()
	fmt.Println(c, "° C is ", k, "° K")
}
