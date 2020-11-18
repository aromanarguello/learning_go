package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
 d := newDeck()

 if len(d) != 16 {
	 t.Errorf("Expected deck length of 16 but got %d", len(d))
 }

 cV := strings.Join([]string(d), ",")[0:3]
 cS := strings.Join([]string(d), ",")[7:13]
 
 if cV != "Ace" {
	 t.Errorf("Expected first cardValue to be Ace but got: %v", cV)
 }

 if cS != "Spades" {
	 t.Errorf("Expected first cardSuite to Be Spades but got: %v", cS)
 }

 if d[0] != "Ace of Spades" {
	 t.Errorf("Expected first card to be Ace of Spades, but got: %v", d[0])
 }

 if d[len(d) - 1] != "Four of Cloves" {
	 t.Errorf("Expected last card to be Four of Cloves, but got: %v", d[len(d) - 1])
 }

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got: %v", len(loadedDeck))

	}

	os.Remove("_decktesting")
}