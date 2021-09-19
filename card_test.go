package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBeat(t *testing.T) {
    for i := 0; i <= 8 ; i++ { 
        card := NewCard(i+1)
        leadCard := NewCard(i)
        trump := NewCard(0)

        want := card.Beat(leadCard, trump)

        assert.Equal(t, true, want,
            "The first card should beat the second one.")
    }
    
    assert.Equal(t, true, NewCard(9).Beat(NewCard(10), NewCard(0)), 
        "The first card should beat the second one.")
}

