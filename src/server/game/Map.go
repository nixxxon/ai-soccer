package game

type MapElement interface {
	Type() string
}

type Map struct {
	Walls []Rectangle	`json:"walls"`
	Goals []Rectangle	`json:"goals"`
}

type Rectangle struct {
	X int		`json:"x"`
	Y int		`json:"y"`
	Width int	`json:"width"`
	Height int	`json:"height"`
}

func DefaultMap() Map {
	walls := []Rectangle{}
	walls = append(walls, Rectangle{-10, -15, -1, 30})
	walls = append(walls, Rectangle{10, 15, 1, -30})
	goals := []Rectangle{}
	goals = append(goals, Rectangle{-2, -15, 4, -1})
	goals = append(goals, Rectangle{2, 15, -4, 1})

	default_map := Map{walls, goals}

	return default_map
}
