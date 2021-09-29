

# BrisGolang

BrisGolang is a Go implementation of the game of [briscola](https://en.wikipedia.org/wiki/Briscola) using the WebSocket protocol for client/server communication.



## Usage

You can play in 2 ways:

1. not using the server
2. using the server



### Not using the server

Start the game with `go run .`

```
$ go run .

------------- brisGolang -------------

game:       
  players:  
    MAX:   [<nil> <nil> <nil>] 0
    MIN:   [<nil> <nil> <nil>] 0
  table:   []
  trump:   

------------- round 1 -------------

game:       
  players:  
    MAX:   [5♥ 6♥ 6♠] 0
    MIN:   [K♣ 4♣ T♥] 0
  table:   []
  trump:   ♥

> cardIdx (0,1,2) [default=0]:
```

At the prompt type the index of the card (`0`, `1`, `2`) you want to play


### Using the server 

Launch the server

```
$ go run server/server.go 
http server started on localhost:8080
```

In another terminal use [wscat](https://github.com/websockets/wscat) (or something similar) to open a connection 

```
$ wscat -c ws://localhost:8080/ws
Connected (press CTRL+C to quit)
> 
```

Type `newGame` to create a new game

```
> newGame
< 
game:       
  players:  
    MAX:   [<nil> <nil> <nil>] 0
    MIN:   [<nil> <nil> <nil>] 0
  table:   []
  trump:   
> 

```

Type `start` to start the game

```
> start
< 
------------- round 1 -------------
< 
game:       
  players:  
    MAX:   [7♠ K♠ J♦] 0
    MIN:   [6♥ T♦ Q♥] 0
  table:   []
  trump:   ♣
> 
```

Type `play` and the index of a card (`0`, `1`, `2`) to play

```
> play 1
< 
game:       
  players:  
    MAX:   [7♠ <nil> J♦] 0
    MIN:   [6♥ T♦ Q♥] 0
  table:   [{MAX K♠}]
  trump:   ♣
< 
game:       
  players:  
    MAX:   [7♠ <nil> J♦] 0
    MIN:   [6♥ T♦ <nil>] 0
  table:   [{MAX K♠} {MIN Q♥}]
  trump:   ♣
< 
------------- round 2 -------------
< 
game:       
  players:  
    MAX:   [7♠ 2♥ J♦] 7
    MIN:   [6♥ T♦ A♣] 0
  table:   []
  trump:   ♣
> 
```


## TODO

* a multiplayer version
* a real client

