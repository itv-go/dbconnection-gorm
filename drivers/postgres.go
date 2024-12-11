package drivers

import (
	"fmt"
	"github.com/itv-go/dbconnection-gorm/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf dbconnectiondto.Config) *gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.HOST,
		conf.PORT,
		conf.USERNAME,
		conf.PASSWORD,
		conf.DATABASE,
	)

	psql := postgres.Open(dsn)
	return &psql
}
