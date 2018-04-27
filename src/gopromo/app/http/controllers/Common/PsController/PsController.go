package PsController

import (
	"gopromo/app/service/PsService"
	"gopromo/utils"

	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//传入参数处理
	r.ParseForm()
	result := PsService.New(r.Form)
	code, url := result.Handle()

	response := utils.ResponseJson{}
	response.Code = code
	response.Data = map[string]string{"url": url}
	response.Show(w)
}
