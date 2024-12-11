package drivers

import (
	"fmt"
	"github.com/itv-go/dbconnection-gorm/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(conf dbconnectiondto.Config) *gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.USERNAME,
		conf.PASSWORD,
		conf.HOST,
		conf.PORT,
		conf.DATABASE,
	)

	msql := mysql.Open(dsn)
	return &msql
}
