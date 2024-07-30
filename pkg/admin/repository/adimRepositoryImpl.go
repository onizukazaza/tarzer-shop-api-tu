package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_adminException "github.com/onizukazaza/tarzer-shop-api-tu/pkg/admin/exception"
)
	

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database , logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{db: db, logger: logger}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)
	if err := r.db.Connect().Create(adminEntity).Scan(admin).Error; err!= nil {
	r.logger.Error("Creating admin failed: %s", err.Error())
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}
	return admin, nil
	}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	if err := r.db.Connect().Where("id =?", adminID).First(admin).Error; err!= nil {
    r.logger.Error("Failed to find admin by ID: %s", err.Error())
        return nil, &_adminException.AdminNotFound{AdminID: adminID}
    }

return admin, nil	
}
