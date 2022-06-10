package main

import (
	"reflect"
	"testing"
)

func Test_newDeck(t *testing.T) {
	d := newDeck()
	expectedDeckSize := 52
	if len(d) != expectedDeckSize {
		t.Errorf("Expected Deck length is %d, but recieved [%d].", expectedDeckSize, len(d))
	}
}

func Test_newDeck_cardByIndex(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  string
	}{
		{"Index 0 should have Ace of Spades", 0, "Ace of Spades"},
		{"Index 12 should have King of Spades", 12, "King of Spades"},
		{"Index 13 should have Ace of Diamonds", 13, "Ace of Diamonds"},
		{"Index 25 should have King of Diamonds", 25, "King of Diamonds"},
		{"Index 26 should have Ace of Hearts", 26, "Ace of Hearts"},
		{"Index 38 should have King of Hearts", 38, "King of Hearts"},
		{"Index 39 should have Ace of Clubs", 39, "Ace of Clubs"},
		{"Index 51 should have King of Clubs", 51, "King of Clubs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeck()[tt.index]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeck() at index [%v] = %v, want %v", tt.index, got, tt.want)
			}
		})
	}
}

func Test_newDeckFromFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDeckFromFile(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDeckFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
