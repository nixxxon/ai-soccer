package network

import (
	"net"
	"fmt"
	"strconv"
	"bufio"
)

type Connection struct {
	conn net.Conn
	command string
}

func NewConnection(greeting string, listener net.Listener) Connection {
	fmt.Print("Awaiting connections:")
	conn, _ := listener.Accept()
	conn.Write([]byte(greeting+"\n"))
	fmt.Print("Found connection")
	connection := Connection{conn, ""}
	go listenToConnection(connection)
	return connection
}

func (connection Connection) SendState(tick int) {
	connection.conn.Write([]byte("Tick "+ strconv.Itoa(tick) +"\n"))
}

func listenToConnection(conn Connection) {
	for {
		message, _ := bufio.NewReader(conn.conn).ReadString('\n')
		conn.command = message
	}

}