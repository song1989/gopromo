/*
	用户信息表
*/
package userTable

import (
	"gopromo/app/models"
	"gopromo/config"
	"gopromo/connect"
)

var usersType []models.User
var userType models.User

type player interface {
	GetData(page int, perPgae int) ([]models.User, error)
	GetById(id int) (models.User, error)
}

type dbType struct {
	*config.MysqlDb
}

func (this *dbType) GetData(page int, perPage int) ([]models.User, error) {
	if this.Err != nil {
		return nil, this.Err
	}
	//this.Db.LogMode(true)
	//defer this.Db.Close()

	offset := (page - 1) * perPage

	//查询
	result := this.Db.Order("id desc").Offset(offset).Limit(perPage).Find(&usersType)
	if sqlErr := result.Error; sqlErr != nil {
		return nil, sqlErr
	}
	return usersType, nil
}

func (this *dbType) GetById(id int) (models.User, error) {
	if this.Err != nil {
		return models.User{}, this.Err
	}
	//db.LogMode(true)
	//defer this.Db.Close()

	result := this.Db.First(&userType, id)
	if sqlErr := result.Error; sqlErr != nil {
		return userType, sqlErr
	}
	return userType, nil
}

func New() player {
	mysql := connect.Mysql{}
	return &dbType{mysql.New()}
}
