package game

import "math"

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Pawn struct {
	Id       int    `json:"id"`
	Position Vector `json:"position"`
	Team     int    `json:"team"`
}

type Ball struct {
	Position Vector `json:"position"`
}

func (pawn *Pawn) canKick(ball *Ball) bool {
	return pawn.Position.distanceTo(ball.Position) < 1
}

func (position Vector) distanceTo(otherPosition Vector) float64 {
	deltaX := position.X - otherPosition.X
	deltaY := position.Y - otherPosition.Y
	return math.Sqrt(deltaX*deltaX + deltaY + deltaY)
}
