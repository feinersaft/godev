package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
)

var oneway bool
var spaceline string
var speed, price int
var w = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

const distance = 62100000

func generateTicket() {
	//w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	speed = rand.Intn(14) + 16
	price = (distance / speed) / 60 / 60 / 24

	if rand.Intn(1) == 1 {
		price *= 2
	}

	// Spaceline
	switch rand.Intn(2) {
	case 0:
		fmt.Fprint(w, "Virgin Galactic")
	case 1:
		fmt.Fprint(w, "SpaceX")
	case 2:
		fmt.Fprint(w, "Space Adventures")
	}

	fmt.Fprint(w, "\t")

	// Days
	fmt.Fprint(w, rand.Intn(30)+1)
	fmt.Fprint(w, "\t")

	// Trip type
	if rand.Intn(1) == 1 {
		fmt.Fprint(w, "One-way")
	} else {
		fmt.Fprint(w, "Round-trip")
		price *= 2
	}
	fmt.Fprint(w, "\t")

	// Price
	fmt.Fprint(w, "$ ")
	fmt.Fprintln(w, int(price))
}

func printTableHeader() {
	fmt.Fprintln(w, "Spaceline\tDays\tTrip type\tPrice")
	fmt.Fprintln(w, "===")
}

func main() {

	printTableHeader()

	for i := 0; i < 10; i++ {
		generateTicket()
	}
	w.Flush()
}
