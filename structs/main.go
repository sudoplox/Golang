package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func main() {
	var alex person
	/* 
	zero value 
	string-> "" 
	int/flaot->0 
	bool->false
	*/
	alex.firstName="Alex"
	alex.lastName="Anderson"
	fmt.Printf("%+v",alex)
}
