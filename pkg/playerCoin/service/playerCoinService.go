package service

import (
	_playerCoinModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/playerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	Showing(playerID string) *_playerCoinModel.PlayerCoinShowing
}
