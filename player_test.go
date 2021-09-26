package main

import (
    //"fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
    player := NewPlayer("MAX", AI{RandomStrategy{}})
    
    assert.Equal(t, "MAX", player.id, "")
    assert.Equal(t, 0, player.score, "")
    assert.Equal(t, [3]*Card{nil, nil, nil}, player.cards, "")
}

func TestIndexOfCard(t *testing.T) {
    player := NewPlayer("MAX", AI{RandomStrategy{}})

    idx := player.indexOfCard(nil)
  
    assert.Equal(t, [3]*Card{nil, nil, nil}, player.cards, "")
    assert.Equal(t, 0, idx, "")
    
    player.cards = [3]*Card{&Card{0,0,0}, nil, nil}
    idx = player.indexOfCard(nil)
    
    assert.Equal(t, 1, idx, "")
    
    player.cards = [3]*Card{&Card{0,0,0}, &Card{0,0,0}, nil}
    idx = player.indexOfCard(nil)
    
    assert.Equal(t, 2, idx, "")
}

func TestTakeCard(t *testing.T) {
    player := NewPlayer("MAX", AI{RandomStrategy{}})
    
    i := 0
    
    for ; i < 3; i++ {
        card := Card{i,0,0}
        idx := player.takeCard(&card)
        
        //fmt.Printf("%v %v %v\n", idx, card, player.cards)
        //fmt.Printf("%v %p %p\n", idx, &card, player.cards[idx])
        
        assert.Equal(t, i, idx, "")
        assert.Equal(t, card, *player.cards[idx], "")
    }
    
    // player has already got 3 cards
    
    cards := player.cards
    card := Card{i,0,0}
    idx := player.takeCard(&card)
    
    //fmt.Printf("%v %v %v %v\n", idx, &card, player.cards[0], cards[0])    
    //fmt.Printf("%v %p %p %p\n", idx, &card, player.cards[0], cards[0])
        
    assert.Equal(t, -1, idx, "")
    assert.Equal(t, cards, player.cards, "")
}

/*
func TestMakeDecision(t *testing.T) {
    player := NewPlayer("MAX")

    idx := player.makeDecision()
    
    fmt.Println(idx)
    
    //assert.Nil(t, card, "")
    //assert.NotEqual(t, -1, idx, "")
}*/

