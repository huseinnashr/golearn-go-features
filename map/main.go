package main

import "fmt"

func main() {
	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ffff",
	// }

	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)
}
