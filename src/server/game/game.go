package game

import (
	"fmt"
	"encoding/json"
)

type Game struct {
	Pawns []Pawn    `json:"pawns"`
	Ball  Ball	`json:"ball"`
	Frame int	`json:"frame"`
	Id    int	`json:"id"`
	Map	Map	`json:"map"`
	TickMs int	`json:"tick_ms"`
}

const NUM_PAWNS int = 8

func CreateGame(gameId int) Game {
	var pawns []Pawn = []Pawn{}
	for i:=0; i< NUM_PAWNS; i++ {
		var position Position
		var team int = i/(NUM_PAWNS/2)
		if( i<4 ) {
			position = Position{float64(-6 + 4*i), -10}
		} else {
			position = Position{float64(-6 + 4*i - 16), +10}
		}
		newPawn := Pawn{Id:i, Position:position, Team:team}
		pawns = append(pawns, newPawn)
	}

	game := Game{pawns, Ball{Position{0, 0}}, 15000, gameId, DefaultMap(), 1000}
	return game
}

func (this *Game) Tick(commands []PawnCommand) {
	//printableState := string(this.ToJsonState())//strconv.Itoa(this.Frame)
	//fmt.Print(printableState + "\n\n")
	fmt.Print("|")
	this.Frame = this.Frame + 1
	//this.Pawns[0].Position.X = this.Pawns[0].Position.X + 2

	for _, command := range commands {
		command.ApplyTo(this)
	}
}

func (this *Game) ToJsonState() []byte {
	jsonState, _ := json.Marshal(this)
	return jsonState
}

func (this *Game) MirrorCopy() *Game {
	mirrorPawns := []Pawn{}
	for _, pawn := range this.Pawns {
		mirrorPos := Position{-pawn.Position.X, -pawn.Position.Y}
		mirrorTeam := (pawn.Team+1)%2
		mirrorId := NUM_PAWNS - 1 - pawn.Id
		mirrorPawn := Pawn{Id:mirrorId, Position:mirrorPos, Team:mirrorTeam}
		mirrorPawns = append([]Pawn{mirrorPawn}, mirrorPawns...)
	}

	mirrorBall := Ball{Position{-this.Ball.Position.X, -this.Ball.Position.Y}}

	return &Game{mirrorPawns, mirrorBall, this.Frame, this.Id, DefaultMap(), 1000}
}
