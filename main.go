package main

import (
    "fmt"
    bg "github.com/calog3r0/brisGolang/game"
)

var gc *bg.GameController

func printHeader(str string) string {
    return fmt.Sprintf("\n------------- %s -------------", str)
}

func printRound() string {
    round := fmt.Sprintf("%s %d", "round", gc.GetGame().GetRound())
    return printHeader(round)
}

func onStateChanged() {
    fmt.Println(gc.GetGame())
    
    if gc.GetGame().IsOver() {
        fmt.Println()
    }
}

func onNewRound() {
    fmt.Println(printRound())
}

func onGameOver() {
    fmt.Println(printHeader("GAME OVER"))
}

func onTakeTurn() {
    controller := gc.GetGame().GetCurrentPlayer().GetController()
    if _, ok := controller.(bg.Human); ok {
        fmt.Print("\n> cardIdx (0,1,2) [default=0]: ")
    }
}

func main() {
    fmt.Println(printHeader("brisGolang"))
    
    //p1 := bg.NewPlayer("MAX", bg.AI{bg.RandomStrategy{}})
    p1 := bg.NewPlayer("MAX", bg.Human{})
    p2 := bg.NewPlayer("MIN", bg.AI{bg.RandomStrategy{}})
    
    gc = bg.NewGameController(p1, p2, onStateChanged, onNewRound, onGameOver,
        onTakeTurn)
    
    //gc.GetGame().SetDeck(gc.GetGame().GetDeck()[1:9])
    
    onStateChanged()

    gc.Start()
}

/*
func main() {
    hFormat := "\n------------- %s -------------\n"
    
    fmt.Printf(hFormat, "brisGolang")

    p1 := bg.NewPlayer("MAX", bg.AI{bg.RandomStrategy{}})
    p2 := bg.NewPlayer("MIN", bg.AI{bg.RandomStrategy{}})

    game := bg.NewGame(p1, p2)

    //game.deck = game.deck[1:9]

    fmt.Println("\ngame:   ", game)
    
    game.Start()
    
    fmt.Printf(hFormat, fmt.Sprintf("%s %d", "round", game.GetRound()))
    
    fmt.Println("\ngame:   ", game)
    
    for ; ! game.IsOver() ; {
    
        if game.GetCurrentPlayer() == p1 {
            fmt.Print("\n> cardIdx (0,1,2) [default=0]: ")
        }
        
        cardIdx := game.GetCurrentPlayer().TakeTurn(game)
            
        if game.Play(cardIdx) {

            fmt.Println("\ngame:   ", game)
            
            if game.RoundCompleted() {
                game.NewRound()
                
                if ! game.IsOver() {
                    fmt.Printf(hFormat, 
                        fmt.Sprintf("%s %d", "round", game.GetRound()))
                } else {
                    fmt.Printf(hFormat, "GAME OVER")
                }
                
                fmt.Println("\ngame:   ", game)
            }
        }
    }

    fmt.Println("")
}
*/

