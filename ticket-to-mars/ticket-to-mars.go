package main

import (
	"fmt"
	"math/rand"
)

var oneway bool
var spaceline string
var speed, price int

const distance = 62100000

func generateTicket() {
	speed = rand.Intn(14) + 16
	price = (distance / speed) / 60 / 60 / 24

	if rand.Intn(1) == 1 {
		price *= 2
	}

	// Spaceline
	switch rand.Intn(2) {
	case 0:
		fmt.Print("Virgin Galactic")
	case 1:
		fmt.Print("SpaceX")
	case 2:
		fmt.Print("Space Adventures")
	}

	fmt.Print("\t")

	// Days
	fmt.Print(rand.Intn(30) + 1)
	fmt.Print("\t")

	// Trip type
	if rand.Intn(1) == 1 {
		fmt.Print("One-way")
	} else {
		fmt.Print("Round-trip")
		price *= 2
	}
	fmt.Print("\t")

	// Price
	fmt.Print("$ ")
	fmt.Println(int(price))
}

func printTableHeader() {
	fmt.Println("Spaceline\tDays\tTrip type\tPrice")
	fmt.Println("=============================================")
}

func main() {

	printTableHeader()

	for i := 0; i < 10; i++ {
		generateTicket()
	}
}
