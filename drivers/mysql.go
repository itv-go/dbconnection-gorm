package drivers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(host string, port int, dbname string, username string, password string) *gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbname,
	)

	msql := mysql.Open(dsn)
	return &msql
}
