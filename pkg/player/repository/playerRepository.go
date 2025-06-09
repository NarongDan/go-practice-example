package repository

import "tutorial/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerId string) (*entities.Player, error)
}
