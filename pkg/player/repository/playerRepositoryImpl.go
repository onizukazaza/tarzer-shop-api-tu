package repository

import ( 
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/databases"
	"github.com/onizukazaza/tarzer-shop-api-tu/entities"
	_playerExceptions "github.com/onizukazaza/tarzer-shop-api-tu/pkg/player/exception"
 )

type PlayerRepositoryImpl struct {
	db databases.Database
	logger echo.Logger

}

func NewPlayerRepositoryImpl (db databases.Database, logger echo.Logger) PlayerRepository {
	return &PlayerRepositoryImpl{db: db, logger: logger}
}

func (r *PlayerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)

	if err := r.db.Connect().Create(playerEntity).Scan(player).Error; err!= nil {
        r.logger.Errorf("Creating player failed: %s", err.Error())
        return nil, &_playerExceptions.PlayerCreating{PlayerID: playerEntity.ID}
    }
	return player, nil
}

func (r *PlayerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	player := new(entities.Player)
	if err := r.db.Connect().Where("id =?", playerID).First(player).Error; err!= nil {
		r.logger.Errorf("Failed to find player by ID: %s", err.Error())
        return nil, &_playerExceptions.PlayerNotFound{PlayerID: playerID}
    }
    return player, nil
}


