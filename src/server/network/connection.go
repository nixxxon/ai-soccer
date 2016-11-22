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

type Connection interface {
	SendState(tick int)
	Disconnect()
	Listen()
	GetCommands() ([]game.PawnCommand)
}

type PlayerConnection struct {
	conn *websocket.Conn
	commands []game.PawnCommand
}

type EmptyConnection struct {

}

func MakeEmptyConnection() Connection {
	fmt.Print("Empty connection created")
	connection := EmptyConnection{}
	//connection.commands = append(connection.commands, game.MoveCommand{Direction:0.5})
	return connection
}

func NewConnection(websocket *websocket.Conn) Connection {
	fmt.Print("Found connection")
	connection := PlayerConnection{websocket, []game.PawnCommand{}}
	//connection.Listen()
	return connection
}

func (connection PlayerConnection) SendState(tick int) {
	connection.conn.Write([]byte("Tick "+ strconv.Itoa(tick) +"\n"))
}

//func listenToConnection(conn Connection) {
//	for {
//		message, _ := bufio.NewReader(conn.conn).ReadString('\n')
//		json.Unmarshal([]byte(message), conn.commands);
//	}
//}

func (this PlayerConnection) Listen() {
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

func (this PlayerConnection) onDisconnect() {
	fmt.Printf("Disconnect!!! BOOOM")
	// TODO remove connection from list
}

func (this PlayerConnection) Disconnect() {
	this.conn.Close()
}

func (this PlayerConnection) GetCommands() ([]game.PawnCommand) {
	return this.commands
}



func (this EmptyConnection) Listen() {}

func (this EmptyConnection) SendState(tick int) {}

func (this EmptyConnection) Disconnect() {}

func (this EmptyConnection) GetCommands() ([]game.PawnCommand) {
	return []game.PawnCommand{}
}
