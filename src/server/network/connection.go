package network

import (
	"fmt"
	"strconv"
	"bufio"
	"./../game"
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
)



type Connection struct {
	conn *websocket.Conn
	commands []game.PawnCommand
}

func MakeDummyConnection() Connection {
	fmt.Print("Empty connection created")
	connection := Connection{nil, []game.PawnCommand{}}
	connection.commands = append(connection.commands, game.MoveCommand{Direction:0.5})
	return connection
}

func NewConnection(websocket *websocket.Conn) Connection {
	fmt.Print("Found connection")
	connection := Connection{websocket, []game.PawnCommand{}}
	//connection.Listen()
	return connection
}

func (connection Connection) SendState(tick int) {
	connection.conn.Write([]byte("Tick "+ strconv.Itoa(tick) +"\n"))
}

//func listenToConnection(conn Connection) {
//	for {
//		message, _ := bufio.NewReader(conn.conn).ReadString('\n')
//		json.Unmarshal([]byte(message), conn.commands);
//	}
//}

func (this Connection) Listen() {
	var in []byte
	for {
		websocket.Message.Send(this.conn, "Na")
		time.Sleep(1*time.Second)

		if err := websocket.Message.Receive(this.conn, &in); err != nil {
			this.onDisconnect()
			return
		}
		fmt.Printf("Batman")
	}
	for {
		message, _ := bufio.NewReader(this.conn).ReadString('\n')
		json.Unmarshal([]byte(message), this.commands);
	}
}

func (this Connection) onDisconnect() {
	fmt.Printf("Disconnect!!! BOOOM")
	// TODO remove connection from list
}

func (this Connection) Disconnect() {
	this.conn.Close()
}
