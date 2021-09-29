package game

import (
    "fmt"
    "errors"
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

func NewGame(player1 *Player, player2 *Player) *Game {
    //player1 := NewPlayer(id1, AI{RandomStrategy{}})
    //player1 := NewPlayer(id1, Human{})
    //player1 := NewPlayer(id1, Remote{})
    //player2 := NewPlayer(id2, AI{RandomStrategy{}})
    
    /*
    players := map[string]*Player{
        id1: &player1,
        id2: &player2,
    }*/
    // https://stackoverflow.com/q/40578646
    // https://stackoverflow.com/q/32751537
    
    players := [2]*Player{player1, player2}

    deck := NewDeck(40)
    deck = deck.shuffle()
    
    return &Game{
        players: players,
        currentPlayer: players[0],
        deck: deck,
        table: NewTable(),
        round: ((40 - len(deck)) / 2) + 1,
    }
}

func (game *Game) GetDeck() Deck {
    return game.deck
}

func (game *Game) SetDeck(deck Deck) {
    game.deck = deck
}

func (game *Game) String() string {
    player1 := game.players[0]
    player2 := game.players[1]
    
    // StringWithPlayersSwapped(swapPlayers bool) string
    /*if swapPlayers {
        player1, player2 = player2, player1
    }*/
    
    var trump string

    if game.trump == nil {  
        trump = ""
    } else {
        trump = suits[game.trump.suit]
    }
    
    return fmt.Sprintf(
        "\ngame:       " +
        "\n  players:  " +
        "\n    %v:   %v" +
        "\n    %v:   %v" +        
        "\n  table:   %v" +
        "\n  trump:   %v", /*+
        "\n  deck:    %v",*/ 
        player1.id, player1,
        player2.id, player2,
        game.table, trump/*, game.deck*/)
}

func (game *Game) GetCurrentPlayer() *Player {
    return game.currentPlayer
}

func (game *Game) GetRound() int {
    return game.round
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

func (game *Game) IsOver() bool {
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

func (game *Game) Start() {
    game.trump = game.deck[len(game.deck)-1]
    
    for _, player := range game.players {
        game.dealCard(player)
        game.dealCard(player)
        game.dealCard(player)
    }
    
    game.currentPlayer.turn, game.nextPlayer().turn = true, false
}

func (game *Game) nextPlayer() *Player {
    if game.currentPlayer == game.players[0] {
        return game.players[1]
    } else {
        return game.players[0]
    }
}
    
func (game *Game) ValidAction(cardIdx int) bool {
    player := game.currentPlayer
    return cardIdx >= 0 && cardIdx <= 2 && player.cards[cardIdx] != nil
}

func (game *Game) Play(cardIdx int) error {
    player := game.currentPlayer
    
    if ! game.ValidAction(cardIdx) {
        return errors.New("cardIdx is not valid")
    }
    
    card := player.play(cardIdx)
    game.table = append(game.table, &Op{player, card})

    game.currentPlayer = game.nextPlayer()
    
    if ! game.RoundCompleted() {
        game.currentPlayer.turn, game.nextPlayer().turn = true, false
    } else {
        game.currentPlayer.turn, game.nextPlayer().turn = false, false
    }

    return nil
}

func (game *Game) RoundCompleted() bool {
    return len(game.table) == len(game.players)
}

func (game *Game) NewRound() {
    if ! game.RoundCompleted() {
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
    
    game.currentPlayer.turn, game.nextPlayer().turn = true, false
    
    game.round++
}

