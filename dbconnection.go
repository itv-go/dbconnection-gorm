package dbconnection

import (
	"errors"
	"github.com/itv-go/dbconnection-gorm/drivers"
	"gorm.io/gorm"
)

func Connect(driver string, host string, port int, dbname string, username string, password string) (*gorm.DB, error) {
	switch driver {
	case "mysql":
		db := drivers.NewPostgres(host, port, dbname, username, password)
		return gorm.Open(*db, &gorm.Config{})
	case "postgres":
		db := drivers.NewPostgres(host, port, dbname, username, password)
		return gorm.Open(*db, &gorm.Config{})
	default:
		return nil, errors.New("database driver not support")
	}
}
