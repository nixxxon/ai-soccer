package network

import "time"
import "./../game"

type Coreographer struct {
	Ai1        Connection
	Ai2        Connection
	Game       game.Game
	Spectators []Connection
}

func RunGame(ai1 Connection, ai2 Connection, game game.Game) {
	coreographer := Coreographer{ai1, ai2, game, []Connection{}}

	for {
		coreographer.Game.Tick(ai1.GetCommands())
		coreographer.postState()

		time.Sleep(100 * time.Millisecond)
	}
}

func (coreo Coreographer) Run() {
	for {
		coreo.Game.Tick(coreo.Ai1.GetCommands())
		coreo.postState()

		time.Sleep(100 * time.Millisecond)
	}
}

func (this Coreographer) postState() {
	this.Ai1.SendState(this.Game.Frame)
	this.Ai2.SendState(this.Game.Frame)

	for _, spectator := range this.Spectators {
		spectator.SendState(this.Game.Frame)
	}
}

func (this Coreographer) AddSpectator(spectator Connection) {
	this.Spectators = append(this.Spectators, spectator)
}

