package main

import "net"
import "fmt"
import (
    "bufio"
    "encoding/json"
    "server/network"
)
import (
    "golang.org/x/net/websocket"
)

func main() {

    // connect to this socket
    conn, err := websocket.Dial("ws://127.0.0.1:3000/lobby", "", "ws://127.0.0.1:3000")
    if err != nil {
        fmt.Print(err)
        fmt.Print("Server not found")
        return
    }
    fmt.Print("Connected")
    handshake, _ := json.Marshal(network.ClientHandshake{GameId:50, Role:"player"})
    response, _ := conn.Write(handshake)
    fmt.Print(response)
    fmt.Print("fdas")

    //printNextMessage(conn)

    var in []byte
    for {
        if err := websocket.Message.Receive(conn, &in); err != nil {
            fmt.Println("Server disconnected by sending bad websocket")
            return
        }
	    fmt.Println(string(in[:]))
        //for {
            	//message, _ := bufio.NewReader(conn).ReadString('\n')
		//fmt.Println(message)
            //json.Unmarshal([]byte(message), commands);
        //}
        // read in input from stdin
        //reader := bufio.NewReader(os.Stdin)
        fmt.Println(" - Internal tick - ")
        //text, _ := reader.ReadString('\n')
        // send to socket
        //fmt.Fprintf(conn, text + "\n")
        //printNextMessage(conn)
    }
}

