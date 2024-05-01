package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_playerCoin "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/exception"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
	"gorm.io/gorm"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin, tx *gorm.DB) (*entities.PlayerCoin, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	playerCoin := new(entities.PlayerCoin)

	if err := conn.Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Error("Player's balance recording failed:", err.Error())
		return nil, &_playerCoin.CoinAdding{}
	}

	return playerCoin, nil
}

func (r *playerCoinRepositoryImpl) Showing(playerID string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoin := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Model(
		&entities.PlayerCoin{},
	).Where(
		"player_id = ?", playerID,
	).Select(
		"player_id, sum(amount) as coin",
	).Group(
		"player_id",
	).Scan(&playerCoin).Error; err != nil {
		r.logger.Error("Calculating player coin failed:", err.Error())
		return nil, &_playerCoin.PlayerCoinShowing{}
	}

	return playerCoin, nil
}
