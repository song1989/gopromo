package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ResponseJson struct {
	Succ bool        `json:"succ"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Time int64       `json:"time"`
}

func (res *ResponseJson) json() string {
	if res.Code == 200 {
		res.Succ = true
	} else {
		res.Succ = false
	}
	if res.Msg == "" {
		res.Msg = http.StatusText(res.Code)
	}
	res.Time = time.Now().Unix()
	webStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return "json err"
	}
	return string(webStr)
}

func (this *ResponseJson) Show(w http.ResponseWriter) {
	fmt.Fprintln(w, this.json())
}
