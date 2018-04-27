package TestController

import (
	"gopromo/app/constant/banyan"
	"gopromo/center/user"
	"gopromo/connect"
	"gopromo/table/user"
	"gopromo/utils"
	"log"
	"net/http"
	"strconv"
)

func Banyan(w http.ResponseWriter, r *http.Request) {
	banyan := connect.Banyan{}
	banyanDb := banyan.New(banyanConstant.NameSpaceInPromo, banyanConstant.TableDeftab, banyanConstant.TypeHset, "test")
	//banyanDb.Set("test", "1111")
	value, _ := banyanDb.GetAll()
	response := utils.ResponseJson{}
	response.Data = value
	response.Code = http.StatusOK
	response.Show(w)
}

func Usertable(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	response := utils.ResponseJson{}
	userTab := userTable.New()

	data := make(map[string]string)
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		user, err := userTab.GetById(id)
		if err != nil {
			data["err2"] = err.Error()
			response.Data = data
			response.Code = http.StatusInternalServerError
			response.Show(w)
			return
		}
		data["name"] = user.Name
		response.Data = data
		response.Code = http.StatusOK
		response.Show(w)
	} else {
		page, perPage := 1, 10
		users, err := userTab.GetData(page, perPage)
		if err != nil {
			data["err3"] = err.Error()
			response.Data = data
			response.Code = http.StatusInternalServerError
			response.Show(w)
			return
		}
		var userList []map[string]interface{}
		for i := 0; i < len(users); i++ {
			item := users[i]
			var uItem = make(map[string]interface{})
			uItem["id"] = item.ID
			uItem["name"] = item.Name
			uItem["age"] = item.Age
			userList = append(userList, uItem)
		}

		response.Data = userList
		response.Code = http.StatusOK
		response.Show(w)
		return
	}
}

func TestOne(w http.ResponseWriter, r *http.Request) {
	//server := utils.URL{}
	//paramMap := server.GetMap(r.URL.Query())
	avatarService := utils.Avatar{}
	avatarUrl, _ := avatarService.Get(2, "/in/2015/06/08/77DD52AF-B9ED-7A6F-8912-DB7313F50D7A.jpg")

	log.Println(avatarUrl)

	response := utils.ResponseJson{}
	response.Code = http.StatusOK
	response.Data = avatarUrl
	response.Show(w)
	return
}

func Testthrift(w http.ResponseWriter, r *http.Request) {
	response := utils.ResponseJson{}
	uc := userCenter.New()
	user, err := uc.GetUserInfoById(1)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Data = map[string]string{"msg": err.Error()}
		response.Show(w)
		return
	}

	response.Code = http.StatusOK
	response.Data = user
	response.Show(w)
	return
}
