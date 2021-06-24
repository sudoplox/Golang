package main

// import (
// 	"fmt"
// )

func main() {

	// var card string = "Ace of Spades"
	
	/*
	var-> variable
	card-> name
	string-> only string type will be assigned ever
	
	available types -> bool, string, int, float64
	
	*/

	// for i,card:=range cards {
	// 	fmt.Println(i,card)
	// }
	
	// /*
	// for index,card := range cards{

	// }
	// index->index of the iter
	// card->iter object
	// */

	cards := newDeck()
	cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

}
