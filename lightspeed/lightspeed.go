package main

import "fmt"

func main() {
	const lightSpeed = 299792
  const hoursPerDay = 24

	var distance = 56000000
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

  distance = 96300000
  var spaceXSpeed = 100800
  fmt.Println(distance/spaceXSpeed/hoursPerDay, "days")
}
