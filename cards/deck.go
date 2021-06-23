package main

// create a new type of deck
// which is a slide of strings
import "fmt"
type deck []string // this new deck type, extends slice of string properties

 

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}
	
	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append( cards, value+" of "+suit)
		}
	}

	return cards
}


// no name -> receiver
// any variable of type deck gets access to print func
// d-> actual copy of deck we are working with, sort of like "this"
// deck->every variable deck can access

func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
	
}

func deal(d deck,handSize int) (deck, deck) {
	return d[ :handSize ],d[ handSize: ]
}

