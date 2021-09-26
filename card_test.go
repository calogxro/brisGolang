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

        want := Beat(card, leadCard, trump)

        assert.Equal(t, true, want,
            "The first card should beat the second one.")
    }

    card := NewCard(9)
    leadCard := NewCard(10)
    trump := NewCard(0)

    assert.Equal(t, true, Beat(card, leadCard, trump), 
        "The first card should beat the second one.")
}

