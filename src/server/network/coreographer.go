package network

import "time"
import (
	"server/game"
	"fmt"
)

type Coreographer struct {
	Ai1        Connection
	Ai2        Connection
	Game       game.Game
	Spectators []Connection
}

func (coreo *Coreographer) Run() {
	for {
		// TODO validate (and sanitize?) commands.
		commands := coreo.Ai1.GetCommands()
		for _, command := range coreo.Ai2.GetCommands() {
			commands = append(commands, command.MirrorCopy())
		}

		coreo.Game.Tick(commands)
		coreo.postState()

		time.Sleep(1000 * time.Millisecond)
	}
}

func (this *Coreographer) postState() {
	this.Ai1.SendState(this.Game.ToJsonState())

	this.Ai2.SendState(this.Game.MirrorCopy().ToJsonState())

	//fmt.Println(strconv.Itoa(len(this.Spectators)))

	for _, spectator := range this.Spectators {
		spectator.SendState(this.Game.MirrorCopy().ToJsonState())
	}
}

func (this *Coreographer) AddSpectator(spectator Connection) {
	this.Spectators = append(this.Spectators, spectator)
	fmt.Printf("added spec")

	//handshake := ServerHandshake{GameId:55, TickMs:1000, WelcomeMessage:"Hello mr spectator to this lovely game!", StaticMap:game.DefaultMap()}
	//jsonHandshake, _ := json.Marshal(handshake)
	//spectator.SendHandshake(jsonHandshake)
}

