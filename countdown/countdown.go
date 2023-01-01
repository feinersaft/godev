// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var seconds = 10

	for seconds > 0 {
		fmt.Printf("%v..\n", seconds)
		time.Sleep(time.Second)
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) == 0 {
			break
		}

		seconds--
	}
	if seconds == 0 {
		fmt.Println("Prost Neujahr!")
	} else {
		fmt.Println("Countdown wegen Böllergeknalle abgebrochen..")
	}
}
