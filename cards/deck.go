package main

// create a new type of deck
// which is a slide of strings
import "fmt"
type deck []string // this new deck type, extends slice of string properties

// no name -> receiver
func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
	
}