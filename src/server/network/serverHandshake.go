package network

import "server/game"

type ServerHandshake struct {
	WelcomeMessage string	`json:"welcome_message"`
	GameId int	`json:"game_id"`
	TickMs int	`json:"tick_ms"`
	StaticMap game.Map `json:"map"`
	Game game.Game	`json:"state"`
}
