package network

import "time"
import "./../game"

type Coreographer struct {
	ai1 Connection
	ai2 Connection
	game game.Game
	spectator Connection
}

func RunGame(ai1 Connection, ai2 Connection, game game.Game) {
	coreographer := Coreographer{ai1, ai2, game, ai1}

	for {
		coreographer.game.Tick(ai1.commands)
		coreographer.postState()

		time.Sleep(100 * time.Millisecond)
	}
}

func (coreo Coreographer) postState() {
	coreo.ai1.SendState(coreo.game.Frame)
	coreo.ai2.SendState(coreo.game.Frame)
}


