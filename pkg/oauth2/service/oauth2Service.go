package service

import (
	_adminModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/admin/model"
	_playerModel "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
}
