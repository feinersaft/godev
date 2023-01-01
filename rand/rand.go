package main

import (
  "fmt"
  "math/rand"
)

func main() {
  var num = rand.Intn(56000000) + 1
  fmt.Println(num)

  num = rand.Intn(401000000) + 1
  fmt.Println(num)
}
