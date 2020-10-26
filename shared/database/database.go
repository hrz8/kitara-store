package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hrz8/kitara-store/shared/config"
)

type (
	// MysqlInterface is represent method in database structs
	MysqlInterface interface {
		Connect() (*gorm.DB, error)
	}

	database struct {
		AppConfig config.AppConfigInterface
	}
)

func (d *database) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		d.AppConfig.GetDbUser(),
		d.AppConfig.GetDbPassword(),
		d.AppConfig.GetDbHost(),
		d.AppConfig.GetDbPort(),
		d.AppConfig.GetDbName(),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewMysql is an factory that implement of mysql database configuration
func NewMysql(appConfig config.AppConfigInterface) MysqlInterface {
	return &database{appConfig}
}
