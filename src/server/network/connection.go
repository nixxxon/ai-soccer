package network

import (
	"net"
	"fmt"
	"strconv"
	"bufio"
	"./../game"
	"encoding/json"
)

type Connection struct {
	conn net.Conn
	commands []game.PawnCommand
}

func EmptyConnection() Connection {
	fmt.Print("Empty connection created")
	connection := Connection{nil, []game.PawnCommand{}}
	connection.commands = append(connection.commands, game.MoveCommand{Direction:0.5})
	return connection
}

func NewConnection(greeting string, listener net.Listener) Connection {
	fmt.Print("Awaiting connections:")
	conn, _ := listener.Accept()
	conn.Write([]byte(greeting+"\n"))
	fmt.Print("Found connection")
	connection := Connection{conn, []game.PawnCommand{}}
	go listenToConnection(connection)
	return connection
}

func (connection Connection) SendState(tick int) {
	if(connection.conn == nil) {
		return;
	}
	connection.conn.Write([]byte("Tick "+ strconv.Itoa(tick) +"\n"))
}

func listenToConnection(conn Connection) {
	for {
		message, _ := bufio.NewReader(conn.conn).ReadString('\n')
		json.Unmarshal([]byte(message), conn.commands);
	}

}