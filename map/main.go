package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// We can also first declare the map
	// var colors map[string]string
	// or use "make" to create a map
	// colors := make(map[string]string)
	// To add a new key/value pair to a map
	// colors["black"] = "#000000"
	// To delete a key/value pair in a map
	// delete("black")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
