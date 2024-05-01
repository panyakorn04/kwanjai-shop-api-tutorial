package repository

import (
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin, tx *gorm.DB) (*entities.PlayerCoin, error)
	Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error)
}
