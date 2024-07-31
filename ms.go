package main

import (
  "fmt"
  "os"
)

func main() {
  mapStr := os.Args[1:]

  mapStr[3]  = "86A0"
  mapStr[11] = "88A0"
  mapStr[12] = mapStr[13]
  mapStr[13] = "88A0"
  mapStr[26] = "83A0"
  mapStr[27] = "85A0"
  mapStr[28] = "84A0"

  fmt.Println(mapStr)
}
