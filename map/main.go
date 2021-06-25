package main

import (
	"fmt"
)


func main() {

	//map[key type]value_type
	var colors1 map[string]string
	//or
	colors2 := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#fbf723",
	// }
	colors2["white"] = "#ff0000"

	delete(colors2,"white") //to delete a pair in map
	fmt.Println(colors1,colors2)
}
