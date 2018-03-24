package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	card := deck{}
	hcard := []string{"Spade", "Heart", "Diamond", "Club"}
	lcard := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, hc := range hcard {
		for _, lc := range lcard {

			card = append(card, lc+" of "+hc)
		}
	}

	return card
}

func (d deck) print() {

	for i, c := range d {
		fmt.Println(i, c)
	}
}

func deal(no int, d deck) (deck, deck) {

	return d[:no], d[no:]
}

func (d deck) saveToFile(fileName string) error {

	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	return err
}

func (d deck) toString() string {

	return strings.Join([]string(d), ", ")
}

func newDeckFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {

	t := time.Now()
	source := rand.NewSource(t.UnixNano())

	r := rand.New(source)

	for i := range d {
		randomPostion := r.Intn(len(d) - 1)

		d[i], d[randomPostion] = d[randomPostion], d[i]
	}
}
