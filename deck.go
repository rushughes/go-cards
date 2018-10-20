package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// a receiver on a function
// essentially lets us extend the type deck with a method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
