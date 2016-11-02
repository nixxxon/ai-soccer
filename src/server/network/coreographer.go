package network

import "time"

type Coreographer struct {
	ai1 Connection
	ai2 Connection
	game int
	spectator Connection
}

func RunGame(ai1 Connection, ai2 Connection, game int) {
	coreographer := Coreographer{ai1, ai2, game, ai1}

	for {
		coreographer.postState()
		coreographer.game = coreographer.game+1
		time.Sleep(100 * time.Millisecond)
	}
}

func (coreo Coreographer) postState() {
	coreo.ai1.SendState(coreo.game)
	coreo.ai2.SendState(coreo.game)
}


