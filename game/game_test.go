package game

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func newGameForTesting() *Game {
    p1 := NewPlayer("MAX", AI{RandomStrategy{}})
    p2 := NewPlayer("MIN", AI{RandomStrategy{}})
    return NewGame(p1, p2)
}

func TestNewGame(t *testing.T) {
    game := newGameForTesting()

    game.deck = game.deck[1:7]

    //fmt.Println("\ngame:   ", game)
    
    assert.Nil(t, game.trump, "")
    
    for _, player := range game.players {
        assert.Equal(t, 3, len(player.cards), "")
        assert.Equal(t, [3]*Card{nil,nil,nil}, player.cards, "")
    }
    
    assert.Equal(t, 6, len(game.deck), "")
}

func TestGame(t *testing.T) {
    //game := newGameForTesting()
    //fmt.Println("game:   ", game)
}

func TestDealCard(t *testing.T) {
    game := newGameForTesting()

    player := game.players[0]
    wantedCard := game.deck[0]
    
    idx := game.dealCard(player)
    
    assert.Equal(t, wantedCard, player.cards[idx], "")
}

func TestGameStart(t *testing.T) {
    game := newGameForTesting()

    game.deck = game.deck[1:9]

    //fmt.Println("\ngame:   ", game)
    
    game.Start()
       
    //fmt.Println("\ngame:   ", game)
    
    assert.Equal(t, 2, len(game.deck), "")
    assert.NotNil(t, game.trump, "")
    assert.Equal(t, game.deck[len(game.deck)-1], game.trump, "")
    
    for _, player := range game.players {
        assert.Equal(t, -1, player.indexOfCard(nil), "")
    }
}

func TestGamePlay(t *testing.T) {
    game := newGameForTesting()

    //fmt.Println("\ngame:   ", game)
    
    game.Start()
    
    length := 0
    
    assert.Equal(t, length, len(game.table), "")
    
    //fmt.Println("\ngame:   ", game)

    for _ = range game.players {
        cardIdx := randomDecision(game)
        
        player := game.currentPlayer
        
        game.Play(cardIdx)
        
        //fmt.Println("\ngame:   ", game)
        
        length++
        
        assert.Equal(t, length, len(game.table), "")
        assert.NotEqual(t, -1, player.indexOfCard(nil), "")
    }
}

func TestNewRound(t *testing.T) {
    game := newGameForTesting()

    fmt.Println("\ngame:   ", game)

    game.Start()
    
    fmt.Println("\ngame:   ", game)
    
    for _ = range game.players {
        cardIdx := randomDecision(game)
        
        game.Play(cardIdx)
        
        fmt.Println("\ngame:   ", game)
    }
    
    nextCards := []*Card{game.deck[0], game.deck[1]}

    game.NewRound()
    
    fmt.Println("\ngame:   ", game)

    assert.NotEqual(t, -1, game.currentPlayer.indexOfCard(nextCards[0]), "")
    assert.NotEqual(t, -1, game.nextPlayer().indexOfCard(nextCards[1]), "")

    assert.Equal(t, 0, len(game.table), "")

    for _, player := range game.players {
        assert.Equal(t, -1, player.indexOfCard(nil), "")
    }
}

