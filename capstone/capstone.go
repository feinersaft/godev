package main

import (
	"fmt"
	"math/rand"
  "time"
)

var spacelines = [3]string{"Space Adv", "SpaceX", "Virgin Galactic"}
const distance = 62100000
var speed, price, duration int
var tripType = [2]string{"Round-trip","One-way"}

func main() {
  fmt.Printf("Spaceline\tDays\tTrip type\tPrice\n")
  fmt.Printf("======================================\n")
	for count := 0; count < 10; count++ {
    rand.Seed(time.Now().UnixNano())
    speed = 16 + rand.Intn(30 - 16)
    price = 36 + rand.Intn(50 - 36)
    duration = distance/speed/86400
		if tripType[rand.Intn(2)] == "Round-trip" {
        price *= 2
    }
    fmt.Printf("%v\t%v\t%v\t$ %v\n", spacelines[rand.Intn(3)], duration, tripType[rand.Intn(2)], price)

	}
}
