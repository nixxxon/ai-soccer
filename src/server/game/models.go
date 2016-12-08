package game

type Position struct {
	X int	`json:"x"`
	Y int	`json:"y"`
}

type Pawn struct {
	Id int	`json:"id"`
	Position Position	`json:"position"`
	Team int		`json:"team"`
}

type Ball struct {
	Position Position	`json:"position"`
}

