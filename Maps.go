package main

import "fmt"

func main() {

	color := map[string]string{

		"Red":   "ff0000",
		"Green": "4bf745",
		"White": "ffffff",
	}

	printMap(color)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}
