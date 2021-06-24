package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { 
	// it all comes in a single test block and not seperate	
	
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got, %v", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs"{
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// need to handle cleanup while testing! delete files if created during test

//wow crazy long name of function
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got, %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}