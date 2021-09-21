package main

import "fmt"

func main() {
	// [string]string means all keys are of type string and all values are of type string
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	/*
	colors["white"] = "#ffffff"
	delete(colors, "white")
	*/

	printMap(colors)
}

// c -> argument name
// map[string]string -> type of the map
func printMap(c map[string]string) {
	// color -> key
	// hex -> value
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}