package main

import (
  "fmt"
  "os"
)

func main() {
  mapStr := os.Args[1:]

  countZeros := 0
  for index, tile := range mapStr {
    if tile == "0" {
      if countZeros == 0 {
        mapStr[index] = "86A0"
      } else if countZeros == 1 {
        mapStr[index] = "88A0"
      }

      countZeros++
    }
  }

  fmt.Println(mapStr)
}
