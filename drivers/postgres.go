package drivers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(host string, port int, dbname string, username string, password string) *gorm.Dialector {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname,
	)

	psql := postgres.Open(dsn)
	return &psql
}
