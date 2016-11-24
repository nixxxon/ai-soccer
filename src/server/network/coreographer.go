package network

import "time"
import (
	"./../game"
	"fmt"
	"strconv"
)

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

		time.Sleep(1000 * time.Millisecond)
	}
}

func (coreo *Coreographer) Run() {
	for {
		coreo.Game.Tick(coreo.Ai1.GetCommands())
		coreo.postState()

		time.Sleep(1000 * time.Millisecond)
	}
}

func (this *Coreographer) postState() {
	this.Ai1.SendState(this.Game.ToJsonState())

	this.Ai2.SendState(this.Game.MirrorCopy().ToJsonState())

	fmt.Print(strconv.Itoa(len(this.Spectators)))

	for _, spectator := range this.Spectators {
		spectator.SendState(this.Game.MirrorCopy().ToJsonState())
	}
}

func (this *Coreographer) AddSpectator(spectator Connection) {
	this.Spectators = append(this.Spectators, spectator)
	fmt.Printf("added spec")
}

