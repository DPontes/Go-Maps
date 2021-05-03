package main

import "fmt"

func main() {

  // Declaring a map with values
  colors := map[string]string {     // this map will have ´keys´ of type string and `values` of type string
    "red"   : "#FF0000",
    "green" : "#0000FF",
    "blue"  : "#00FF00",
  }

  printMap(colors)
}

func printMap(c map[string]string) {
  for color, hex := range c {
    fmt.Println("hex code for", color, "is", hex)
  }
}
