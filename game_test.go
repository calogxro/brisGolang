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
    
    assert.Equal(t, 6, len(*game.deck), "")
}

func TestGame(t *testing.T) {
    //game := NewGame("MAX", "MIN", NewDeck(6))
    //fmt.Println("game:   ", game)
}

func TestDealCard(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))    

    player := game.players["MAX"]
    deck := *game.deck

    lastCard := &deck[len(deck)-1]
    
    idx := game.dealCard(player)
    
    assert.Equal(t, lastCard, player.cards[idx], "")
}

func TestGameStart(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))

    //fmt.Println("\ngame:   ", game)
    
    game.start()
       
    //fmt.Println("\ngame:   ", game)
    
    assert.NotNil(t, game.trump, "")
    
    for _, player := range game.players {
        assert.Equal(t, -1, player.indexOfCard(nil), "")
    }
    
    assert.Equal(t, 0, len(*game.deck), "")
}

func TestGamePlay(t *testing.T) {
    game := NewGame("MAX", "MIN", NewDeck(6))

    //fmt.Println("\ngame:   ", game)
    
    game.start()
    
    length := 0
    
    assert.Equal(t, length, len(game.table), "")
    
    //fmt.Println("\ngame:   ", game)

    for _, player := range game.players {
        cardIdx := player.makeDecision()
        
        game.play(player, cardIdx)
        
        //fmt.Println("\ngame:   ", game)
        
        length++
        
        assert.Equal(t, length, len(game.table), "")
        assert.NotEqual(t, -1, player.indexOfCard(nil), "")
    }
    
    assert.Equal(t, len(game.players), length, "")
}

func TestNewRound(t *testing.T) {
    deck := NewDeck(40)
    deck = deck.shuffle()
    deck = deck[1:7]
    
    game := NewGame("MAX", "MIN", deck)

    fmt.Println("\ngame:   ", game)
    
    game.start()
    
    fmt.Println("\ngame:   ", game)
    
    for _, player := range game.players {
        cardIdx := player.makeDecision()
        
        game.play(player, cardIdx)
        
        fmt.Println("\ngame:   ", game)
    }
    
    game.newRound()
    
    fmt.Println("\ngame:   ", game)
}

