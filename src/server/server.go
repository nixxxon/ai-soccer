package main

import "net"
import "fmt"
//import "bufio"
//import "strings" // only needed below for sample processing
import "./network"

func main() {

    fmt.Println("Launching server...")

    // listen on all interfaces
    ln, _ := net.Listen("tcp", ":8081")

    conn1 := network.NewConnection("Welcome player 1! ", ln)
    fmt.Println("One player joined")

    conn2 := network.NewConnection("Welcome player 2! ", ln)
    fmt.Println("Two players joined")

    go network.RunGame(conn1, conn2, 0)


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
