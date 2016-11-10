package game


type PawnCommand interface {
	ApplyTo(pawn *Pawn)
	Type() string
}

type MoveCommand struct {
	Direction float32
}

func (command MoveCommand) ApplyTo(pawn *Pawn) {
	pawn.Position.X = pawn.Position.X + 1
}

func (command MoveCommand) Type() string {
	return "move"
}
