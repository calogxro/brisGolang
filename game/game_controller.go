package game

import (
    //"fmt"
)

type GameController struct {
    game *Game
    onStateChanged func()
    onNewRound func()
    onGameOver func()
    onTakeTurn func()
}

func NewGameController(p1 *Player, p2 *Player, onStateChanged func(), 
    onNewRound func(), onGameOver func(), onTakeTurn func()) *GameController {
    return &GameController{
        game: NewGame(p1, p2), 
        onStateChanged: onStateChanged, 
        onNewRound: onNewRound,
        onGameOver: onGameOver,
        onTakeTurn: onTakeTurn,
    }
}

func (gc *GameController) GetGame() *Game {
    return gc.game
}

func (gc *GameController) Start() {
    gc.game.Start()
    gc.onNewRound()
    gc.onStateChanged()
    gc.loop()
}

// called by players' controllers
func (gc *GameController) Play(cardIdx int) {
    gc.game.Play(cardIdx)
    gc.onStateChanged()
    
    if gc.game.RoundCompleted() {  
        gc.game.NewRound()
        if ! gc.game.IsOver() {
            gc.onNewRound()
        } else {
            gc.onGameOver()
        }
        gc.onStateChanged()
    }
    
    gc.loop()
}

func (gc *GameController) loop() {
    if ! gc.game.IsOver() {
        if ! gc.game.RoundCompleted() {  
            player := gc.game.GetCurrentPlayer()
            gc.onTakeTurn()
            player.TakeTurn(gc.game, gc.Play)
        }
    }
}

