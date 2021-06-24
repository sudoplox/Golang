package main

// create a new type of deck
// which is a slide of strings
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string  {
	return strings.Join([]string(d),",")	
}
func(d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(
		filename, 
		[]byte(d.toString()),
		0666)
}

func newDeckFromFile(filename string) deck {
	bs,err := ioutil.ReadFile(filename)
	//err-> value of type error, if all good then nil
	if err != nil{
		// option1: log the error and call newDeck
		// option2: log the error and quit program 👍
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split( string(bs), "," ) //Ace of Spades,Two o....
	return deck(ss)
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano()) // pull in the current time
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn( len(d)-1 ) 
		d[i],d[newPosition]=d[newPosition],d[i]
	}
}