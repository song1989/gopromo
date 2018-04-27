/**
  包含banyan连接池，空间、表明常量定义
*/
package connect

import (
	"gopromo/app/constant/banyan"
	"gopromo/env"
	"gopromo/libraries/banyan"
	"log"
	"strings"
)

type Banyan struct {
}

type player interface {
	GetDb() *banyan_api.BanyanClient
}

type banyanTable struct {
	Np    string
	Table string
}

func (this *banyanTable) GetDb() *banyan_api.BanyanClient {
	conf := strings.Split(env.BanyanService, ",")
	cluster := banyan_api.NewClusterClient(conf)
	db, err := cluster.GetBanyanClient(this.Np, this.Table, 3000, 3)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func (this *Banyan) NewSource(np string, table string) player {
	return &banyanTable{Np: np, Table: table}
}

func (this *Banyan) New(np string, table string, mold string, key string) (result banyan_api.UseFunc) {
	banyanDb := this.NewSource(np, table)
	switch mold {
	case banyanConstant.TypeSet:
		result = banyan_api.NewBanyanUse(banyan_api.OperationSet{banyanDb.GetDb()})
	case banyanConstant.TypeHset:
		result = banyan_api.NewBanyanUse(banyan_api.OperationHset{banyanDb.GetDb(), key})
	case banyanConstant.TypeZset:
		result = banyan_api.NewBanyanUse(banyan_api.OperationZset{banyanDb.GetDb(), key})
	}
	return result
}
