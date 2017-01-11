package game

type MapElement interface {
	Type() string
}

type Map struct {
	Walls []Rectangle `json:"walls"`
	Goals []Rectangle `json:"goals"`
}

type Rectangle struct {
	Position Vector `json:"position"`
	Size     Vector `json:"size"`
}

func DefaultMap() Map {

	walls := []Rectangle{}
	walls = append(walls, Rectangle{Position: Vector{X: -10, Y: -15}, Size: Vector{X: -1, Y: 30}})
	walls = append(walls, Rectangle{Position: Vector{X: 10, Y: 15}, Size: Vector{X: 1, Y: -30}})
	walls = append(walls, Rectangle{Position: Vector{X: -11, Y: -15}, Size: Vector{X: 9, Y: -1}})
	walls = append(walls, Rectangle{Position: Vector{X: 11, Y: -15}, Size: Vector{X: -9, Y: -1}})
	walls = append(walls, Rectangle{Position: Vector{X: -11, Y: 15}, Size: Vector{X: 9, Y: 1}})
	walls = append(walls, Rectangle{Position: Vector{X: 11, Y: 15}, Size: Vector{X: -9, Y: 1}})

	goals := []Rectangle{}
	goals = append(goals, Rectangle{Position: Vector{X: -2, Y: -15}, Size: Vector{X: 4, Y: -1}})
	goals = append(goals, Rectangle{Position: Vector{X: 2, Y: 15}, Size: Vector{X: -4, Y: 1}})

	default_map := Map{walls, goals}

	return default_map
}
