package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/databases"
	"github.com/panyakorn04/kwanjai-shop-api-tutorial/entities"
	_adminException "github.com/panyakorn04/kwanjai-shop-api-tutorial/pkg/admin/exception"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err != nil {
		r.logger.Errorf("Create player error: %v", err)
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id = ?", adminID).First(admin).Error; err != nil {
		r.logger.Errorf("Find player by ID error: %v", err)
		return nil, &_adminException.AdminNotFound{AdminID: admin.ID}
	}

	return admin, nil
}
