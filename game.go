package main

import (
    "fmt"
)

type Op struct {
    player *Player
    card *Card
}

func (op Op) String() string {
    return fmt.Sprintf("{%s %s}", op.player.id, op.card)
}

type Table []*Op    // order matters

func NewTable() Table {
    return []*Op{}
}

type Game struct {
    //players map[string]*Player
    players [2]*Player
    currentPlayer *Player
    deck Deck
    trump *Card
    table Table
    round int
}

func NewGame(id1 string, id2 string, deck Deck) Game {
    //player1 := NewPlayer(id1, AI{RandomStrategy{}})
    player1 := NewPlayer(id1, Human{})
    player2 := NewPlayer(id2, AI{RandomStrategy{}})
    
    /*
    players := map[string]*Player{
        id1: &player1,
        id2: &player2,
    }*/
    // https://stackoverflow.com/q/40578646
    // https://stackoverflow.com/q/32751537
    
    players := [2]*Player{&player1, &player2}

    deck = deck.shuffle()
    
    return Game{
        players: players,
        currentPlayer: players[0],
        deck: deck,
        table: NewTable(),
        round: ((40 - len(deck)) / 2) + 1,
    }
}

func (game Game) String() string {
    player1 := game.players[0]
    player2 := game.players[1]

    var trump string

    if game.trump == nil {
        trump = ""
    } else {
        trump = suits[game.trump.suit]
    }
        
    return fmt.Sprintf(
        "\n  players:   " +
        "\n    %v:   %v %v" +
        "\n    %v:   %v %v" +        
        "\n  table:   %v" +
        "\n  trump:   %v",
        //"\n  deck:    %v", 
        player1.id, player1.cards, player1.score, 
        player2.id, player2.cards, player2.score,
        game.table, trump)//, game.deck)
}

func (game *Game) actions() []int {
    //return game.currentPlayer.cards

    actions := []int{}

    for cardIdx, card := range game.currentPlayer.cards {
        if card != nil {
            actions = append(actions, cardIdx)
        }
    }

    return actions
}

func (game *Game) isOver() bool {
    p1 := game.players[0]
    p2 := game.players[1]
    
    return 0 == len(game.deck) + len(game.table) + p1.nCards + p2.nCards
}

func (game *Game) dealCard(player *Player) int {
    card, deck := game.deck.pop()
    game.deck = deck
    cardIdx := player.takeCard(card)
    return cardIdx
}

func (game *Game) start() {
    game.trump = game.deck[len(game.deck)-1]
    
    for _, player := range game.players {
        game.dealCard(player)
        game.dealCard(player)
        game.dealCard(player)
    }
}

func (game *Game) nextPlayer() *Player {
    if game.currentPlayer == game.players[0] {
        return game.players[1]
    } else {
        return game.players[0]
    }
}
    
func (game *Game) play(cardIdx int) bool {
    player := game.currentPlayer
    
    if cardIdx >= 0 && cardIdx <= 2 && player.cards[cardIdx] != nil {
        card := player.play(cardIdx)
        game.table = append(game.table, &Op{player, card})

        game.currentPlayer = game.nextPlayer()
  
        return true
    }

    return false
}

func (game *Game) roundCompleted() bool {
    return len(game.table) == len(game.players)
}

func (game *Game) newRound() {
    if ! game.roundCompleted() {
        return
    }
    
    leadOp := game.table[0]
    lastOp := game.table[1]
    
    var winner *Player
    
    if Beat(lastOp.card, leadOp.card, game.trump) {
        winner = lastOp.player
    } else {
        winner = leadOp.player
    }
    
    winner.score += lastOp.card.score + leadOp.card.score

    game.currentPlayer = winner

    game.table = NewTable()

    if len(game.deck) >= len(game.players) {
        game.dealCard(game.currentPlayer)
        game.dealCard(game.nextPlayer())
    }
    
    game.round++
}

