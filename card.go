package main

var suits = [4]string{"♥", "♦", "♣", "♠"}
var ranks = [10]string{"2", "4", "5", "6", "7", "J", "Q", "K", "T", "A"}
var scores = [10]int{0, 0, 0, 0, 0, 2, 3, 4, 10, 11}

type Card struct {
    rank int    // a number in [0, 9]
    suit int    // a number in [0, 3]
    score int
}

func NewCard(idx int) Card {
    return Card{
        rank: idx % 10,
        suit: idx / 10,
        score: scores[idx % 10],
    }
}

func (card Card) String() string {
    return ranks[card.rank] + suits[card.suit]
}

func (card Card) Beat(lead_card Card, trump Card) bool {
    if card.suit == lead_card.suit {
        return card.rank > lead_card.rank
    } else {
        if card.suit == trump.suit {
            return true
        } else {
            return false
        }
    }
}

