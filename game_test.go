package main

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))

    //fmt.Println("\ngame:   ", game)
    
    assert.Nil(t, game.trump, "")
    
    for _, player := range game.players {
        assert.Equal(t, 3, len(player.cards), "")
        assert.Equal(t, [3]*Card{nil,nil,nil}, player.cards, "")
    }
    
    assert.Equal(t, 6, len(game.deck), "")
}

func TestGame(t *testing.T) {
    //game := NewGame("MAX", "MIN", NewDeck(6))
    //fmt.Println("game:   ", game)
}

func TestDealCard(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))    

    player := game.players[0]
    wantedCard := game.deck[0]
    
    idx := game.dealCard(player)
    
    assert.Equal(t, wantedCard, player.cards[idx], "")
}

func TestGameStart(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(8))

    //fmt.Println("\ngame:   ", game)
    
    game.start()
       
    //fmt.Println("\ngame:   ", game)
    
    assert.Equal(t, 2, len(game.deck), "")
    assert.NotNil(t, game.trump, "")
    assert.Equal(t, game.deck[len(game.deck)-1], game.trump, "")
    
    for _, player := range game.players {
        assert.Equal(t, -1, player.indexOfCard(nil), "")
    }
}

func TestGamePlay(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))

    //fmt.Println("\ngame:   ", game)
    
    game.start()
    
    length := 0
    
    assert.Equal(t, length, len(game.table), "")
    
    //fmt.Println("\ngame:   ", game)

    for _ = range game.players {
        cardIdx := randomDecision(game)
        
        player := game.currentPlayer
        
        game.play(cardIdx)
        
        //fmt.Println("\ngame:   ", game)
        
        length++
        
        assert.Equal(t, length, len(game.table), "")
        assert.NotEqual(t, -1, player.indexOfCard(nil), "")
    }
}

func TestNewRound(t *testing.T) {
    deck := NewDeck(40)
    deck = deck.shuffle()
    deck = deck[1:9]
    
    game := NewGame("MAX", "MIN", deck)

    fmt.Println("\ngame:   ", game)
    
    game.start()
    
    fmt.Println("\ngame:   ", game)
    
    for _ = range game.players {
        cardIdx := randomDecision(game)
        
        game.play(cardIdx)
        
        fmt.Println("\ngame:   ", game)
    }
    
    nextCards := []*Card{game.deck[0], game.deck[1]}

    game.newRound()
    
    fmt.Println("\ngame:   ", game)

    assert.NotEqual(t, -1, game.currentPlayer.indexOfCard(nextCards[0]), "")
    assert.NotEqual(t, -1, game.nextPlayer().indexOfCard(nextCards[1]), "")

    assert.Equal(t, 0, len(game.table), "")

    for _, player := range game.players {
        assert.Equal(t, -1, player.indexOfCard(nil), "")
    }
}

