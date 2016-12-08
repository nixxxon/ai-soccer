package game

import "math"

type Position struct {
	X float64	`json:"x"`
	Y float64	`json:"y"`
}

type Pawn struct {
	Id int	`json:"id"`
	Position Position	`json:"position"`
	Team int		`json:"team"`
}

type Ball struct {
	Position Position	`json:"position"`
}

func (pawn *Pawn) canKick(ball *Ball) bool {
	return pawn.Position.distanceTo(ball.Position) < 1
}

func (position Position) distanceTo(otherPosition Position) float64 {
	deltaX := position.X - otherPosition.X
	deltaY := position.Y - otherPosition.Y
	return math.Sqrt(deltaX*deltaX + deltaY+deltaY)
}
