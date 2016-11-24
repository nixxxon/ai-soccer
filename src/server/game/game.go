package game

import (
    "fmt"
    "encoding/json"
)

type Game struct {
    Pawns []Pawn    `json:"doods"`
    Frame int
    Id    int
}

func CreateGame() Game {
    players := []Pawn{Pawn{Position{1, 1}}, Pawn{Position{3, 3}}}

    game := Game{players, 15000, 55}
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
        mirrorPawn := Pawn{mirrorPos}
        mirrorPawns = append([]Pawn{mirrorPawn}, mirrorPawns...)
    }

    return &Game{mirrorPawns, this.Frame, this.Id}
}
