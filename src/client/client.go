package main

import "net"
import "fmt"
import "bufio"

func printNextMessage(conn net.Conn) {
    // listen for reply
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err == nil {
        fmt.Print("Message from server: "+message)

    }
}

func main() {

    // connect to this socket
    conn, err := net.Dial("tcp", "127.0.0.1:8081")
    if err != nil {
        fmt.Print("Server not found")
        return
    }


    printNextMessage(conn)
    for {
        // read in input from stdin
        //reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        //text, _ := reader.ReadString('\n')
        // send to socket
        //fmt.Fprintf(conn, text + "\n")
        printNextMessage(conn)
    }
}

