package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_playerException "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/player/exception"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepository(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err != nil {
		r.logger.Errorf("Create player error: %v", err)
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Where("id = ?", playerID).First(player).Error; err != nil {
		r.logger.Errorf("Find player by ID error: %v", err)
		return nil, &_playerException.PlayerNotFound{PlayerID: playerID}
	}

	return player, nil
}
