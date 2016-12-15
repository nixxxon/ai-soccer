package network

type ClientHandshake struct {
	Role string	`json:"role"`
	GameId int	`json:"game_id"`
}
