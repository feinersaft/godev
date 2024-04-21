package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func generateDates() {
	//rand.Seed(time.Now().UnixNano())
	year := rand.Intn(2023) + 1
	//year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
		if year%4 == 0 {
			daysInMonth++
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}

func main() {
	for count := 0; count < 10; count++ {
		generateDates()
	}
}
