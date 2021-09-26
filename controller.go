package main

import (
    "fmt"
    "math/rand"
)

type Controller interface {
    takeTurn(Game, func(int))
}

/********* Human *********/

type Human struct {
}

func (controller Human) takeTurn(game Game, onAction func(int)) {
    idx, err := readFromStdIn()

    if err == nil {
    }
    
    onAction(idx)
}

func readFromStdIn() (int, error) {
    var i int

    _, err := fmt.Scanf("%d", &i)
    
    if err != nil {
        //fmt.Println("ERROR:", err)
    }  
    
    return i, err
}

/********* AI *********/

type Strategy interface {
    makeDecision(Game) int
}

type RandomStrategy struct {
}

type AI struct {
    strategy Strategy
}

func (strategy RandomStrategy) makeDecision(game Game) int {
    return randomDecision(game)
}

func randomDecision(game Game) int {
    actions := game.actions()
    idx := rand.Intn(len(actions))
    return actions[idx]
}

func (ai AI) takeTurn(game Game, onAction func(int)) {
    action := ai.strategy.makeDecision(game)
    onAction(action)
}

