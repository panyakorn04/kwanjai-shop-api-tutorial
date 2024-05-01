package service

import (
	entities "github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
	_playerCoinRepository "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/repository"
)

type playerCoinService struct {
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository _playerCoinRepository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinService{
		playerCoinRepository: playerCoinRepository,
	}
}

func (s *playerCoinService) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoin, err := s.playerCoinRepository.CoinAdding(playerCoinEntity, nil)
	if err != nil {
		return nil, err
	}

	return playerCoin.ToPlayerCoinMold(), nil
}

func (s *playerCoinService) Showing(playerID string) *_playerCoinModel.PlayerCoinShowing {
	coin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &_playerCoinModel.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}

	return coin
}
