package game

import (
    "fmt"
    "strconv"
)

type Game struct {
    Players []Player
    Frame int
}

func CreateGame() Game {
    players := []Player{Player{Position{1, 1}}, Player{Position{3, 3}}}

    game := Game{players, 15000}
    return game
}

func (this *Game) Tick() {
    printableState := strconv.Itoa(this.Frame)
    fmt.Print(printableState + "\n")
    this.Frame = this.Frame + 1
}
