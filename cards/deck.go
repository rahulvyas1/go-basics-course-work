package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//deck type which is slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}

	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile() {
	str := d.toString()
	message := []byte(str)
	err := ioutil.WriteFile("hello.deck", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (d deck) toString() string {
	s := []string(d)
	var joined string

	joined = strings.Join(s, ", ")

	return joined
}

func newDeckFromFile(f string) deck {
	dat, err := ioutil.ReadFile(f)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// check(err)
	// fmt.Println(string(dat))
	return deck(strings.Split(string(dat), ","))
}

func (d deck) shuffle() deck {
	t := d
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len([]string(t)), func(i, j int) { t[i], t[j] = t[j], t[i] })
	return deck(t)
}
