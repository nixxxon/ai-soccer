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
}

const NUM_PAWNS int = 8

func CreateGame() Game {
	var pawns []Pawn = []Pawn{}
	for i:=0; i< NUM_PAWNS; i++ {
		var position Position
		var team int = i/(NUM_PAWNS/2)
		if( i<4 ) {
			position = Position{-6 + 4*i, -10}
		} else {
			position = Position{-6 + 4*i, +10}
		}
		newPawn := Pawn{Id:i, Position:position, Team:team}
		pawns = append(pawns, newPawn)
	}

	game := Game{pawns, 15000, 55}
	return game
}

func (this *Game) Tick(commands []PawnCommand) {
	//printableState := string(this.ToJsonState())//strconv.Itoa(this.Frame)
	//fmt.Print(printableState + "\n\n")
	fmt.Print("|")
	this.Frame = this.Frame + 1
	this.Pawns[0].Position.X = this.Pawns[0].Position.X + 2
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
		mirrorId := NUM_PAWNS - pawn.Id
		mirrorPawn := Pawn{Id:mirrorId, Position:mirrorPos, Team:mirrorTeam}
		mirrorPawns = append([]Pawn{mirrorPawn}, mirrorPawns...)
	}

	return &Game{mirrorPawns, this.Frame, this.Id}
}
