package config

import (
	"math/rand"

	"gopromo/env"
)

type serverList map[int]string

var imgDomain = map[int]serverList{
	1: serverList{
		1: env.ImgU1,
		2: env.ImgU2,
		3: env.ImgU3,
		4: env.ImgU4,
		5: env.ImgU5,
		6: env.ImgU6,
	},
	2: serverList{
		1: env.ImgIn1,
		2: env.ImgIn2,
		3: env.ImgIn3,
		4: env.ImgIn4,
		5: env.ImgIn5,
		6: env.ImgIn6,
	},
	3: serverList{
		1: env.ImgDown1,
		2: env.ImgDown2,
		3: env.ImgDown3,
	},
	4: serverList{
		1: env.ImgWd1,
		2: env.ImgWd2,
		3: env.ImgWd3,
		4: env.ImgWd4,
		5: env.ImgWd5,
		6: env.ImgWd6,
	},
	5: serverList{
		1: env.ImgBchat,
	},
}

type ImgServer struct {
}

func (this *ImgServer) Get(server int) string {
	serverList, ok := imgDomain[server]
	if ok == false {
		return ""
	}
	serverLen := len(serverList)
	serverRand := rand.Intn(serverLen)
	serverStr, _ := serverList[serverRand]
	return serverStr
}
