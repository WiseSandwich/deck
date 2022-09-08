package main

import "testing"

func testNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 2000 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

}
