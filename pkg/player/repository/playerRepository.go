package repository

import "github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
