package utils

import (
	"github.com/speps/go-hashids"
)

func conf() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = "D@A2(F8*6~"

	hashId, _ := hashids.NewWithData(hd)
	return hashId
}

type Encryption struct {
}

func (h *Encryption) Encode(idInt []int) string {
	hash := conf()
	resultId, _ := hash.Encode(idInt)
	return resultId
}

func (h *Encryption) Decode(idStr string) []int {
	hash := conf()
	resultId := hash.Decode(idStr)
	return resultId
}
