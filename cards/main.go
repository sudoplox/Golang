package main

// import "fmt"

func main() {

	// var card string = "Ace of Spades"
	
	/*
	var-> variable
	card-> name
	string-> only string type will be assigned ever
	
	go is statically typed language like c++, java

	available types -> bool, string, int, float64
	
	*/
	
	// card := "Ace of Spades"
	// := for NEW VARIABLE, initialisation

	// card = "Five of Diamonds"
	// = for REASSGINMENT


	// card := newCard()
	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds",newCard()} //a slice of type string	
	// cards = append(cards, "Six of Spades") //does not modify, returns new
	
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
	
	cards.print()

}
