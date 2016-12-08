package game

type Position struct {
	X int
	Y int
}

type Pawn struct {
	Id int
	Position Position
	Team int
}

type Ball struct {
	Position Position
}

