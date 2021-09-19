package main

import (
    "fmt"
    "math/rand"
)

type Player struct {
    id string
    score int
    cards [3]*Card
}

func NewPlayer(id string) Player {
    return Player{
        id: id,
    }
}

func (player *Player) String() string {
    return fmt.Sprintf(
        "%v %v %v", player.id, player.cards, player.score)
}

func (player *Player) indexOfCard(card *Card) int {
   for k, v := range player.cards {
       if card == v {
           return k
       }
   }
   return -1    //not found.
}
// https://stackoverflow.com/q/8307478

func (player *Player) takeCard(card *Card) int {
    idx := player.indexOfCard(nil)
    if idx != -1 {
        player.cards[idx] = card
    }
    return idx
}

func (player *Player) makeDecision() int {
    //if card == nil {}
    idx := rand.Intn(3)
    return idx
}

