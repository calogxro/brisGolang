package game

import (
    "fmt"
)

type Player struct {
    id string
    score int
    nCards int
    cards [3]*Card
    controller Controller
    turn bool
}

func NewPlayer(id string, c Controller) *Player {
    return &Player{
        id: id,
        controller: c,
    }
}

func (player *Player) String() string {
    /*var turn string
    if player.turn {
        turn = "*"
    } else {
        turn = " "
    }*/
    return fmt.Sprintf(
        "%v %v", player.cards, player.score)
}

func (player *Player) GetId() string {
    return player.id
}

func (player *Player) SetId(id string) {
    player.id = id
}

func (player *Player) GetController() Controller {
    return player.controller
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
        player.nCards++
    }
    return idx
}

func (player *Player) TakeTurn(game *Game, onAction func(int)) {
    player.controller.takeTurn(game, onAction)
}

func (player *Player) play(cardIdx int) *Card {
    card := player.cards[cardIdx]
    player.cards[cardIdx] = nil
    player.nCards--
    
    return card
}
