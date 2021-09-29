package main

import (
    "errors"
    "flag"
    "fmt"    
	"log"
	"net/http"
	"strconv"
	"strings"	
	
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
    
    bg "github.com/calog3r0/brisGolang/game"
)

var id int

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

/*var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}*/
/*
// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
*/
var clients = make(map[*websocket.Conn]bool) // connected clients
//var broadcast = make(chan Message)           // broadcast channel
var broadcast = make(chan []byte)           // broadcast channel
//var broadcast = make(chan string)           // broadcast channel

func main() {
    flag.Parse()
	log.SetFlags(0)
    
    r := mux.NewRouter()
    r.HandleFunc("/ws", handleConnections)
    //r.HandleFunc("/ws/{playerId}", handleConnections)
	http.Handle("/", r)
	
	// Start listening for messages to broadcast
	go handleBroadcast()

    log.Println("http server started on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// on client connection
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	
	id++
	
	//vars := mux.Vars(r)
	playerId := "MAX" //vars["playerId"] //strconv.Itoa(id)
	
	// Create a player to be paired with this connection
	player := bg.NewPlayer(playerId, bg.Remote{})
	
	// Register our new client
	clients[ws] = true

	for {
		//var msg Message
		// Read in a new message as JSON and map it to a Message object
		//err := ws.ReadJSON(&msg)
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		handleCommand(string(msg), player)
	}
}

func getCommand(words []string) (cmd string, params []string) {
    length := len(words)
    if length >= 1 {
        cmd = words[0]
    }
    if length >= 2 {
        params = words[1:length]
    }
    return cmd, params
}

func handleCommand(str string, player *bg.Player) (state string) {
	//var validCommand bool = true
    
    cmd, params := getCommand(strings.Fields(str))
    
    var err error
    
    /*
    if player.GetId() == "" && cmd != "userid" {
        return "userid is required"
    }*/
    
    switch cmd {
        //case "userid": err = userid(params, player)
	    case "newGame": err = newGame(params, player)
	    case "start": err = start(params, player)
	    case "play": err = play(params, player)
	    default: //validCommand = false
	}
    
	if err != nil {	
	    log.Println(err)
	}
	
	return 
}

func handleBroadcast() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
		    err := client.WriteMessage(1, msg)
    		//err := client.WriteJSON(game)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

var p1 *bg.Player
var p2 *bg.Player
var gc *bg.GameController

func newGame(params []string, player *bg.Player) error {
    //p1 = bg.NewPlayer("MAX", bg.Remote{})
    p1 = player
    p2 = bg.NewPlayer("MIN", bg.AI{bg.RandomStrategy{}})
    
    gc = bg.NewGameController(p1, p2, onStateChanged, onNewRound, onGameOver,
        onTakeTurn)
    
    gc.GetGame().SetDeck(gc.GetGame().GetDeck()[1:9])
    
    onStateChanged()
    return nil
}

func start(params []string, player *bg.Player) error {
    if gc == nil {
        return errors.New("error: game not initialized")
    }
    gc.Start()
    return nil
}

func printHeader(str string) string {
    return fmt.Sprintf("\n------------- %s -------------", str)
}

func printRound() string {
    round := fmt.Sprintf("%s %d", "round", gc.GetGame().GetRound())
    return printHeader(round)
}

func onStateChanged() {
    state := gc.GetGame().String()
	broadcast <- []byte(state)
}

func onNewRound() {
    str := printRound()
    broadcast <- []byte(str)
}

func onGameOver() {
    str := printHeader("GAME OVER")
    broadcast <- []byte(str)
}

func onTakeTurn() {
}

func play(params []string, connPlayer *bg.Player) error {
    if len(params) == 0 {
        return errors.New("cardIdx is required")
    }
    
    cardIdx, err := strconv.Atoi(params[0])
    if err != nil {
        return err
    } 
    
    validPlayer := connPlayer == gc.GetGame().GetCurrentPlayer()
    
    if ! validPlayer {
        return errors.New("it's not your turn")
    }
    
    gc.Play(cardIdx)
    
    return nil
}

/*
func userid(params []string, player *bg.Player) error {
    if len(params) == 0 {
        return errors.New("userid is required")
    }
    
    player.SetId(params[0])
    
    return nil
}*/

