package repository

import "github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Admin) (*entities.Admin, error)
	FindByID(adminID string) (*entities.Admin, error)
}
