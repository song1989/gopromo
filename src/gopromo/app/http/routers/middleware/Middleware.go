package middleware

import (
	"gopromo/center/user"
	"gopromo/config"
	"gopromo/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Midd struct {
}

//日志记录所有请求及其路径和处理所需的时间。
func (this Midd) Logging() Middleware {
	//创建新的中间件
	return func(f http.HandlerFunc) http.HandlerFunc {
		//定义 http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			//做中间件的东西
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, r.Method, time.Since(start))
			}()
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

//判断请求方式
func (this Midd) Method(method string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				response := utils.ResponseJson{}
				response.Code = http.StatusBadRequest
				response.Show(w)
				return
			}
			f(w, r)
		}

	}
}

//判断是否登录
func (this Midd) IsLogin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, isOk := "", false

			tgAuthName := "tg_auth"
			token = r.URL.Query().Get("_token")
			if token == "" {
				tgAuth, _ := r.Cookie(tgAuthName)
				token = tgAuth.Value
			}
			if token != "" {
				uc := userCenter.New()
				userField, err := uc.GetUserInfoByPrivateKey(token)
				if err == nil && userField.Id != "" {
					tgAuthValue := userField.PrivateKey
					//查询到用户之后，设置cookie
					cookie := http.Cookie{Name: tgAuthName, Value: tgAuthValue, Path: "/", MaxAge: 86400}
					http.SetCookie(w, &cookie)
					isOk = true
				}
			}
			if isOk == false {
				response := utils.ResponseJson{}
				response.Code = http.StatusUnauthorized
				response.Show(w)
				return
			}
			f(w, r)
		}
	}
}

//判断promo_id是否存在
func (this Midd) PromoId() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			isOk := false
			promoId := r.Form.Get("promo_id")
			if promoId != "" {
				encr := utils.Encryption{}
				promoIdList := encr.Decode(promoId)
				promoIdInt := promoIdList[0]
				if promoIdInt != 0 {
					value := config.Promo{promoIdInt}
					isOk = value.IssetId()
					if isOk {
						r.Form.Set("promo_id", strconv.Itoa(promoIdInt))
					}
				}
			}

			if isOk == false {
				response := utils.ResponseJson{}
				response.Code = http.StatusBadRequest
				response.Data = map[string]string{"error": "not promo"}
				response.Show(w)
				return
			}
			f(w, r)
		}
	}
}

func (this Midd) Validate(rules map[string]string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			val := validate{}
			response := val.logic(r, rules)
			if response.Code != http.StatusOK {
				response.Show(w)
				return
			}
			f(w, r)
		}
	}
}

func (this Midd) Chain(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}
