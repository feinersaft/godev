package main

import (
  "math/rand"
  "fmt"
  "log"
  "time"
)

func main() {
  fmt.Println("Please enter a number (1-100):")
  var pickedNumber, diff string
  n, err := fmt.Scanln(&pickedNumber)
  
  if err != nil {
    log.Fatal(err)
  }
  
  var guess int
  
  for pickedNumber != guess {
    rand.Seed(time.Now(),UnixNano())
    guess = rand.Intn(100)
    if pickedNumber > guess {
      diff = "small"
    } else {
      diff = "big"
    }
    fmt.Printf("My guess was %v and is too %v...\n", guess, diff)
  }
  
  fmt.Printf("I guess your number is %v!\n", guess)
}
