package main

import (
  "fmt"
  "os"
)

func main() {
  mapStr := os.Args[1:]

  for index, tile := range mapStr {
    if tile == "0" {
      fmt.Println("Found a zero! at", index)
    }
  }

  fmt.Println(mapStr)
}
