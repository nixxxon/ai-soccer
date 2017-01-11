package network

import (
	"fmt"
	"bufio"
	"server/game"
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
	"math"
	"math/rand"
)

type Connection interface {
	SendState([]byte)
	SendHandshake([]byte)
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

func (connection PlayerConnection) SendState(state []byte) {
	connection.conn.Write(state)
	fmt.Println("Sent state to player.")
	//fmt.Println(string(state))
}

func (connection PlayerConnection) SendHandshake(handshake []byte) {
	connection.conn.Write(handshake)
	fmt.Println("Sending handshake:")
	fmt.Println(string(handshake))
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
		time.Sleep(1*time.Second)

		if err := websocket.Message.Receive(this.conn, &in); err != nil {
			this.onDisconnect()
			fmt.Println("Client disconnected by sending bad websocket")
			return
		}
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

func (this EmptyConnection) SendState(state []byte) {}

func (this EmptyConnection) SendHandshake(handshake []byte) {}

func (this EmptyConnection) Disconnect() {}

func (this EmptyConnection) GetCommands() ([]game.PawnCommand) {
	commands := []game.PawnCommand{}
	commands = append(commands, game.MoveCommand{Direction:-math.Pi/4, PawnId:0})
	commands = append(commands, game.MoveCommand{Direction:0, PawnId:1})
	commands = append(commands, game.MoveCommand{Direction:math.Pi, PawnId:2})

	if( rand.Intn(3) != 0) {
		commands = append(commands, game.MoveCommand{Direction:math.Pi/2 + math.Pi*0.18, PawnId:3})
	}

	commands = append(commands, game.KickBallCommand{Direction:math.Pi/2, PawnId:3, Force:4})

	return commands
}
