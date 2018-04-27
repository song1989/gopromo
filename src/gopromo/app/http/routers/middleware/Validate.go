package middleware

import (
	"gopromo/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type validate struct {
}

func (this *validate) logic(r *http.Request, rules map[string]string) (response utils.ResponseJson) {
	response = utils.ResponseJson{Code: http.StatusOK}
	for field, rule := range rules {
		if rule != "" {
			value := r.URL.Query().Get(field)
			ruleList := strings.Split(rule, "|")
			for k := 0; k < len(ruleList); k++ {
				if this.check(value, ruleList[k]) == false {
					response.Code = http.StatusBadRequest
					response.Msg = "params " + field + " error"
					return
				}
			}
		}
	}
	return
}

func (this *validate) check(value string, rule string) bool {
	ruleList := strings.Split(rule, ":")
	switch ruleList[0] {
	case "required":
		if value == "" {
			return false
		}
	case "between":
		valueInt, _ := strconv.Atoi(value)
		if valueInt != 0 {
			section := strings.Split(ruleList[1], ",")
			startNum, _ := strconv.Atoi(section[0])
			endNum, _ := strconv.Atoi(section[1])
			if startNum > valueInt || valueInt > endNum {
				return false
			}
		}
	case "url":
		if strings.Contains(value, "http://") == false && strings.Contains(value, "https://") == false {
			return false
		}
	case "int":
		if _, intErr := strconv.Atoi(value); intErr != nil {
			return false
		}
	case "mobile":
		reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
		match, _ := regexp.MatchString(reg, value)
		if match == false {
			return false
		}
	}
	return true
}
