package main

// import "net"
import "fmt"
import "log"
import "net/http"
import "os"
import "golang.org/x/net/websocket"
import (
    "path/filepath"
    "server/game"
    "server/network"
    "encoding/json"
    "strconv"
)
//import "bufio"
//import "strings" // only needed below for sample processing

 //import (
 //    "fmt"
 //    "code.google.com/p/go.net/websocket"
 //    "net/http"
 //)

// func webHandler(ws *websocket.Conn) {
//     var in []byte
//     if err := websocket.Message.Receive(ws, &in); err != nil {
//         return
//     }
//     fmt.Printf("Received: %s\n", string(in))
//     websocket.Message.Send(ws, in)
// }

var connections []network.Connection
//var awaitingSpectators []network.Connection
var coreographers []*network.Coreographer

func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    startDummyGame()

    fmt.Println("Launching servers..." + (filepath.Join(dir, "static")))

    fmt.Println("Launching servers...")

    http.Handle("/", http.FileServer(http.Dir(filepath.Join(dir, "static"))))
    go http.ListenAndServe(":8000", nil)

    fmt.Println("Listening...")

    http.Handle("/lobby", websocket.Handler(webHandler))
    go http.ListenAndServe(":3000", nil)

     //err := http.ListenAndServe(":3000", nil)
     //if err != nil {
     //    panic("ListenAndServe: " + err.Error())
     //}

    //emptyGame := game.CreateGame()
    //go network.RunGame(network.EmptyConnection(), network.EmptyConnection(), emptyGame)

    // // listen on all interfaces
    // ln, _ := net.Listen("tcp", ":8081")

    // conn1 := network.NewConnection("Welcome player 1! ", ln)
    // fmt.Println("One player joined")

    // conn2 := network.NewConnection("Welcome player 2! ", ln)
    // fmt.Println("Two players joined")

    // networkGame := game.CreateGame()
    // go network.RunGame(conn1, conn2, networkGame)

    // fmt.Print(networkGame)

    // run loop forever (or until ctrl-c)
    i := 0
    for i < 100{
        // will listen for message to process ending in newline (\n)
        //message, _ := bufio.NewReader(conn1).ReadString('\n')
        // output message received
        //fmt.Print("Message Received:", string(message))
        // sample process for string received
        // newmessage := strings.ToUpper(message)

        //conn1.SendState(i)

        // send new string back to client
        // conn1.Write([]byte(newmessage + "\n"))
        // conn2.Write([]byte(newmessage + "\n"))
    }
}

func startDummyGame() {
    fmt.Println("Starting dummy game")
    conn1 := network.MakeEmptyConnection()
    conn2 := network.MakeEmptyConnection()
    delegateNewConnection(conn1, 55)
    delegateNewConnection(conn2, 55)
}

func webHandler(ws *websocket.Conn) {
    connection := network.NewConnection(ws)

    var firstMessage []byte
    if err := websocket.Message.Receive(ws, &firstMessage); err != nil {
        return
    }

    fmt.Println(string(firstMessage))
    handshake := network.ClientHandshake{}
    json.Unmarshal(firstMessage, &handshake)

    if(handshake.Role == "spectator") {
        delegateNewSpectator(connection, handshake.GameId)
    } else {
        delegateNewConnection(connection, handshake.GameId)
        if handshake.GameId >= 10000 {
            fmt.Println("Adding an extra dummy opponent")
            conn2 := network.MakeEmptyConnection()
            delegateNewConnection(conn2, handshake.GameId)
        }
    }

    connection.Listen()
}
func delegateNewSpectator(connection network.Connection, requestedGameId int) {
    fmt.Println("Welcome spectator!")
    for _, coreographer := range coreographers {
        if( coreographer.Game.Id == requestedGameId ) {
            coreographer.AddSpectator(connection)
            fmt.Println("Spectating game "+strconv.Itoa(requestedGameId))
            //go coreographer.Run()
            //go coreographer.Run()
            return
        }
        fmt.Println("Searching")
    }
    fmt.Println("game "+strconv.Itoa(requestedGameId)+" not found")
    connection.Disconnect()
}

func delegateNewConnection(connection network.Connection, requestedGameId int) {
    fmt.Println("Welcome Player!")
    connections = append(connections, connection)
    if (len(connections) >= 2) {
        startNewGame(requestedGameId)
    }
}

func startNewGame(gameId int) {
    networkGame := game.CreateGame(gameId)

    conn1 := connections[0]
    conn2 := connections[1]
    connections = connections[2:]

    coreographer := network.Coreographer{conn1, conn2, networkGame, nil}
    go coreographer.Run()

    coreographers = append(coreographers, &coreographer)
    //go network.RunGame(connections[0], connections[1], networkGame)
}
