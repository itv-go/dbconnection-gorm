package dbconnection

import (
	"errors"
	"github.com/itv-go/dbconnection-gorm/drivers"
	"github.com/itv-go/dbconnection-gorm/dto"
	"gorm.io/gorm"
)

func Connect(driver string, config dbconnectiondto.Config) (*gorm.DB, error) {
	switch driver {
	case "mysql":
		db := drivers.NewPostgres(config)
		return gorm.Open(*db, &gorm.Config{})
	case "postgres":
		db := drivers.NewPostgres(config)
		return gorm.Open(*db, &gorm.Config{})
	default:
		return nil, errors.New("database driver not support")
	}
}
