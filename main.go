package main

import ( 
    "fmt"
)

func main() {

    hFormat := "\n------------- %s -------------\n"
    
    fmt.Printf(hFormat, "brisGolang")

    deck := NewDeck(40)
    deck = deck.shuffle()
    //deck = deck[1:9]

    game := NewGame("MAX", "MIN", deck)

    fmt.Println("\ngame:   ", game)
    
    fmt.Printf(hFormat, fmt.Sprintf("%s %d", "round", game.round))

    game.start()

    fmt.Println("\ngame:   ", game)
    
    for ; ! game.isOver() ; {
    
        if game.currentPlayer.id == "MAX" {
            fmt.Print("\n> cardIdx (0,1,2) [default=0]: ")
        }
        
        game.currentPlayer.takeTurn(game, func(cardIdx int){
            
            if game.play(cardIdx) {

                fmt.Println("\ngame:   ", game)
                
                if game.roundCompleted() {
                    game.newRound()
                    
                    if ! game.isOver() {
                        fmt.Printf(hFormat, 
                            fmt.Sprintf("%s %d", "round", game.round))
                    } else {
                        fmt.Printf(hFormat, "GAME OVER")
                    }
                    
                    fmt.Println("\ngame:   ", game)
                }
            }
        })
    }

    fmt.Println("")
}

