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
    players map[string]*Player
    currentPlayer string
    deck *Deck
    trump *Card
    table Table
    round int
}

func NewGame(id1 string, id2 string, deck Deck) Game {
    player1 := NewPlayer(id1)
    player2 := NewPlayer(id2)
    
    players := map[string]*Player{
        id1: &player1,
        id2: &player2,
    }
    // https://stackoverflow.com/q/40578646
    // https://stackoverflow.com/q/32751537
    
    deck = deck.shuffle()
    
    return Game{
        players: players,
        deck: &deck,
        table: NewTable(),
    }
}

func (game Game) String() string {
    player1 := game.players["MAX"]
    player2 := game.players["MIN"]
        
    return fmt.Sprintf(
        "\n  players:   " +
        "\n    %v:   %v %v" +
        "\n    %v:   %v %v" +        
        "\n  table:   %v" +
        "\n  trump:   %v" +
        "\n  deck:    %v", 
        player1.id, player1.cards, player1.score, 
        player2.id, player2.cards, player2.score,
        game.table, game.trump, *game.deck)
}

func (game Game) dealCard(player *Player) int {
    card := game.deck.pop()
    idx := player.takeCard(&card)
    return idx
}

func (game *Game) start() {
    deck := *game.deck
    game.trump = &deck[0]
    
    for _, player := range game.players {
        game.dealCard(player)
        game.dealCard(player)
        game.dealCard(player)
    } 
}

func (game *Game) play(player *Player, cardIdx int) {
    card := player.cards[cardIdx]
    player.cards[cardIdx] = nil
    game.table = append(game.table, &Op{player, card})
}

func (game *Game) newRound() {
    if len(game.table) != len(game.players) {
        return
    }
    
    leadOp := game.table[0]
    lastOp := game.table[1]
    
    score := lastOp.card.score + leadOp.card.score
    
    if lastOp.card.Beat(*leadOp.card, *game.trump) {
        lastOp.player.score += score
    } else {
        leadOp.player.score += score
    }
    
    game.table = NewTable()
}

