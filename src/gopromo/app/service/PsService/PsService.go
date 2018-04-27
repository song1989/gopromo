/*
	图片处理
*/
package PsService

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	configPromo "gopromo/config"
	"gopromo/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func New(params url.Values) Handleer {
	return &handleParams{params}
}

type Handleer interface {
	Handle() (int, string) //图片处理
}

type handleParams struct {
	Params url.Values
}

func (this *handleParams) Handle() (int, string) {
	params := this.Params
	promoId := params.Get("promo_id")

	promoIdInt, err := strconv.Atoi(promoId)
	if err != nil {
		log.Errorln("err:", err)
		return http.StatusBadRequest, ""
	}

	promoConf := configPromo.Promo{promoIdInt}
	psName := promoConf.GetPhotoHandleName()

	log.Errorln("ps name:", psName)
	confValue, _ := getConfigValue(psName)
	//拼接处理图片链接
	getServerUser := func() (string, bool) {
		urlArr := []string{"http://", confValue.Host, "/", confValue.Name, "?"}

		url := strings.Join(urlArr, "")
		for j := 0; j < len(confValue.Field); j++ {
			field := confValue.Field[j]
			param := params.Get(field)
			if param == "" {
				return "", true
			}
			var str []string
			if j == 0 {
				str = []string{field, "=", param}
			} else {
				str = []string{"&", field, "=", param}
			}
			url += strings.Join(str, "")
		}
		return url, false
	}
	serverUrl, fieldErr := getServerUser()
	if fieldErr == true {
		log.Errorln("(handle photo handle fail) err:", fieldErr)
		return http.StatusBadRequest, ""
	}

	log.Println("serverUrl:", serverUrl)
	httpService := utils.Http{}
	result, getErr := httpService.GetContent(serverUrl)
	if getErr != nil {
		log.Errorln("(handle photo handle fail) err:", getErr)
		return http.StatusBadGateway, ""
	}

	type handleRequest struct {
		Code     int    `json:"code"`
		ImageUrl string `json:"image_url"`
		Msg      string `json:"msg"`
	}
	var requsetValue handleRequest
	json.Unmarshal([]byte(result), &requsetValue)

	if requsetValue.ImageUrl == "" {
		log.Errorln("(handle photo json fail) err:", getErr)
		return http.StatusBadGateway, ""
	}
	return http.StatusOK, requsetValue.ImageUrl
}
