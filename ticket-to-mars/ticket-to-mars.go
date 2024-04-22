package main

import (
	"fmt"
	"math/rand"
)

var oneway bool
var spaceline string
var speed int
var price float32

const distance = 62100000

func generateTicket() (string, int) {
	speed = rand.Intn(14) + 16
	//seconds := distance / speed
	//spaceline := rand.Intn(2)

	if rand.Intn(1) == 1 {
		price *= 2
	}
	return spaceline, int(price)
}

func printTable() {
	fmt.Println("Spaceline\tDays\tTrip type\tPrice")
	fmt.Println("=============================================")
}

func main() {

	printTable()

	fmt.Print("$ ")
	fmt.Print(speed)
}
