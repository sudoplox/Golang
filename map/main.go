package main

import (
	"fmt"
)


func main() {

	//map[key type]value_type

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#fbf723",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func  printMap(c map[string]string){
	for col,hex := range c{
		fmt.Println("Hex code for",col,"is:",hex)
	}
}
