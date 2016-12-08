package game

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

