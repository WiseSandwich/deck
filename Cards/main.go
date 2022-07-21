package main

import "WiseSandwich.com/deck/Cards/subpackage"

func main() {
	cards := newDeckFromFile("my")

	cards.print()
	subpackage.PublicFunction("Ricardo")
}
