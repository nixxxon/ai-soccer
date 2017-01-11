package game

import "math"

const runspeed float64 = 1.0

type PawnCommand interface {
	ApplyTo(game *Game)
	Type() string
	GetPawnId() int
	MirrorCopy() PawnCommand
}

type MoveCommand struct {
	// All directions are in radians
	PawnId int
	Direction float64
}

func (command MoveCommand) ApplyTo(game *Game) {
	pawn := &game.Pawns[command.PawnId]
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

type KickBallCommand struct {
	// All directions are in radians
	PawnId int
	Direction float64
	Force float64
}

func (command KickBallCommand) ApplyTo(game *Game) {
	ball := &game.Ball
	pawn := &game.Pawns[command.PawnId]
	if( pawn.canKick(ball)) {
		ball.Position.X = ball.Position.X + command.Force*math.Cos(command.Direction)
		ball.Position.Y = ball.Position.Y + command.Force*math.Sin(command.Direction)
	}
}

func (command KickBallCommand) Type() string {
	return "kick_ball"
}

func (command KickBallCommand) GetPawnId() int {
	return command.PawnId
}

func (command KickBallCommand) MirrorCopy() PawnCommand {
	mirrorPawnId := NUM_PAWNS - 1 - command.PawnId
	mirrorDirection := command.Direction + math.Pi
	return KickBallCommand{mirrorPawnId, mirrorDirection, command.Force}
}
