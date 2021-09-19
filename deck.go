package main

import (
    //"fmt"
    "math/rand"
    "time"
)

type Deck []Card

func NewDeck(numOfCards int) Deck {
    deck := Deck(make([]Card, numOfCards))

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

// ma serve prendere dalla fine?
// altrimenti puoi: "return deck[0], deck[1:]" (aw4y)
func (pDeck *Deck) pop() Card {
    deck := *pDeck
    *pDeck = deck[:len(deck)-1]
    return deck[len(deck)-1]
}
// https://play.golang.org/p/jg_yfpdNXoM

func pop(deck Deck) (Card, Deck) {
    return deck[len(deck)-1], deck[:len(deck)-1]
}

