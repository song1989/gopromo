package config

import "github.com/jinzhu/gorm"

type MysqlDb struct {
	Db  *gorm.DB
	Err error
}
