package game

import "math"

const runspeed float64 = 0.2

type PawnCommand interface {
	ApplyTo(pawn *Pawn)
	Type() string
	GetPawnId() int
	MirrorCopy() PawnCommand
}

type MoveCommand struct {
	// All directions are in radians
	PawnId int
	Direction float64
}

func (command MoveCommand) ApplyTo(pawn *Pawn) {
	pawn.Position.X = pawn.Position.X + runspeed*math.Cos(command.Direction)
	pawn.Position.Y = pawn.Position.Y + runspeed*math.Sin(command.Direction)
}

func (command MoveCommand) Type() string {
	return "move"
}

func (command MoveCommand) GetPawnId() int {
	return command.PawnId
}

func (command MoveCommand) MirrorCopy() PawnCommand {
	mirrorPawnId := NUM_PAWNS - 1 - command.PawnId
	mirrorDirection := command.Direction + math.Pi
	return MoveCommand{mirrorPawnId, mirrorDirection}
}
