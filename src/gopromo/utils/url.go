package utils

import (
	"net/url"
)

type URL struct {
}

func (this *URL) GetMap(values url.Values) map[string]string {
	var paramMap = make(map[string]string)
	for k, v := range values {
		param := v[0]
		if param != "" {
			paramMap[string(k)] = string(param)
		}
	}
	return paramMap
}
