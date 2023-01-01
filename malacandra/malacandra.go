package main

import "fmt"

func main() {
  const hoursPerDay = 24
  var (
    speed int
    distance = 56000000
    days = 28
  )

  speed = distance/(days*hoursPerDay)

  fmt.Printf("A ship would need to travel with a speed of %v km/h\n", speed)
}
