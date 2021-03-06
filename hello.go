package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello")

  x := 5
  y := 7
  sum := x + y
  fmt.Println(sum)

  a := []int{5, 4, 3, 2, 1}
  a = append(a, 13)
  fmt.Println(a)

  vertices := make(map[string]int)
  vertices["triangle"] = 2
  vertices["square"] = 3
  vertices["dodecagon"] = 12

  delete(vertices, "square")
  fmt.Println(vertices)
}
