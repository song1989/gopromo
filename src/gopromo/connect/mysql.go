package connect

import (
	"gopromo/config"
	"gopromo/env"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlInstance *config.MysqlDb
var once sync.Once

type Mysql struct {
}

func (this *Mysql) New() *config.MysqlDb {
	once.Do(func() {
		type conf struct {
			Ip       string
			Port     int
			User     string
			Password string
			Dbname   string
		}
		myConf := conf{
			Ip:       env.MysqlSaveIp,
			Port:     env.MysqlSavePort,
			User:     env.MysqlSaveUser,
			Password: env.MysqlSavePassword,
			Dbname:   env.MysqlSaveDbName,
		}
		var strArr = []string{myConf.User, ":", myConf.Password, "@tcp(", myConf.Ip, ":", strconv.Itoa(myConf.Port), ")/", myConf.Dbname, "?", "charset=utf8&parseTime=True&loc=Local"}
		str := strings.Join(strArr, "")
		log.Println(str)

		db, err := gorm.Open("mysqlConnect", str)
		db.LogMode(true)
		if err != nil {
			log.Println("mysqlConnect connect err:", err)
		}
		mysqlInstance = &config.MysqlDb{db, err}
	})
	return mysqlInstance
}
