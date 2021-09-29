package game

import (
    //"fmt"
    "math/rand"
    "time"
)

type Deck []*Card

// numOfCards is a number in [0...40]
func NewDeck(numOfCards int) Deck {
    deck := Deck(make([]*Card, numOfCards))

    for i, _ := range deck {
        deck[i] = NewCard(i)
    }

    return deck
}

func (deck Deck) shuffle() Deck {
    rand.Seed(time.Now().Unix())

    rand.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })
    
    return deck
}

func (deck Deck) pop() (*Card, Deck) {
    return deck[0], deck[1:]
}
// https://play.golang.org/p/jg_yfpdNXoM
// https://stackoverflow.com/q/38013922
// https://go101.org/article/unofficial-faq.html#error-non-name

func pop(deck Deck) (*Card, Deck) {
    return deck[0], deck[1:]
}

