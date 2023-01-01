package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Please enter a number (1-100):")
	var pickedNumber, guess int
	fmt.Scanln(&pickedNumber)

	for pickedNumber != guess {
		rand.Seed(time.Now().UnixNano())
		guess = rand.Intn(100)
		if pickedNumber > guess {
			fmt.Printf("My guess was %v and is too small...\n", guess)
		} else {
			fmt.Printf("My guess was %v and is too big...\n", guess)
		}

	}

	fmt.Printf("I guess your number is %v!\n", guess)
}
