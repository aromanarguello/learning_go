package main


import "fmt"

func main() {

	colors := map[string]string{
		"red": "#f0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, ":", value)
	}
}