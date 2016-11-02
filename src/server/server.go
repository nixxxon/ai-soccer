package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing
// import "server/game"

func main() {

    fmt.Println("Launching server...")

    // listen on all interfaces
    ln, _ := net.Listen("tcp", ":8081")

    // accept connection on port
    conn, _ := ln.Accept()
    conn.Write([]byte("Welcome player 1!\n"))
    conn2, _ := ln.Accept()
    conn2.Write([]byte("Welcome player 2!\n"))

    // game := Game{[
    //     Player{Position{1, 1}}
    // ]}
    // fmt.Print(game)

    // run loop forever (or until ctrl-c)
    for {
        // will listen for message to process ending in newline (\n)
        message, _ := bufio.NewReader(conn).ReadString('\n')
        // output message received
        fmt.Print("Message Received:", string(message))
        // sample process for string received
        newmessage := strings.ToUpper(message)
        // send new string back to client
        conn.Write([]byte(newmessage + "\n"))
        conn2.Write([]byte(newmessage + "\n"))
    }
}
