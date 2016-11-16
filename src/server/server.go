package main

// import "net"
import "fmt"
import "log"
import "net/http"
import "os"
import "golang.org/x/net/websocket"
import (
    "path/filepath"
    //"./network"
    //"./game"
    "time"
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


func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Launching servers..." + (filepath.Join(dir, "static")))

    fmt.Println("Launching servers...")

    http.Handle("/", http.FileServer(http.Dir(filepath.Join(dir, "static"))))
    go http.ListenAndServe(":80", nil)

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

func webHandler(ws *websocket.Conn) {
    var in []byte
    if err := websocket.Message.Receive(ws, &in); err != nil {
        return
    }
    fmt.Printf("Received: %s\n", string(in))

    fmt.Printf("Second message: %s\n", string(in))
    websocket.Message.Send(ws, in)
    bombard(ws)
}

func bombard(ws *websocket.Conn) {
    var in []byte
    for {
        websocket.Message.Send(ws, "Na")
        time.Sleep(1*time.Second)

        if err := websocket.Message.Receive(ws, &in); err != nil {
            fmt.Printf("Disconnect")
            return
        }
        fmt.Printf("Batman")
    }
}
