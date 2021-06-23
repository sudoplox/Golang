package main

// create a new type of deck
// which is a slide of strings
import "fmt"
type deck []string // this new deck type, extends slice of string properties

// no name -> receiver
// any variable of type deck gets access to print func
// d-> actual copy of deck we are working with, sort of like "this"
// deck->every variable deck can access

func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
	
}