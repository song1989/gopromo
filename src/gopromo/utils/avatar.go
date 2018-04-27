package utils

import (
	"errors"
	"gopromo/config"
	"strconv"
	"strings"
)

type Avatar struct {
	Wh      string
	Size    int
	Suffix  string
	IsHttp  bool
	NetName string
}

func (this *Avatar) Get(server int, url string) (string, error) {
	if url == "" {
		return "", errors.New("not url")
	}
	if string([]byte(url)[:4]) == "http" {
		return url, nil
	}
	if this.Suffix == "" {
		this.Suffix = "jpg"
	}

	//处理server，新图片走无锡，老图片走嘉兴
	urlArr := strings.Split(url, "/")
	urlTime := urlArr[2] + "-" + urlArr[3] + "-" + urlArr[4]
	if server == 1 && urlArr[1] == "in" && urlTime >= "2014-11-24" {
		server = 4
	}
	if server == 4 && urlArr[1] == "in" && urlTime < "2014-11-24" {
		server = 1
	}

	//此处域名，固定域名，避免客户端端对同一张图片因为域名变化而重复加载
	imgService := config.ImgServer{}
	domain := imgService.Get(server)
	if this.IsHttp == true {
		domain = "https://" + domain
	} else {
		domain = "http://" + domain
	}

	var imageUrl string
	switch server {
	case 2:
		if this.Size == 0 {
			this.Wh = "h"
			this.Size = 640
		}
		var compress string
		if this.Wh == "w" {
			compress = strconv.Itoa(this.Size) + "x%3E"
		} else {
			compress = "x" + strconv.Itoa(this.Size) + "%3E"
		}

		var net string
		if this.NetName == "wifi" {
			net = "90"
		} else {
			net = "80"
		}
		imageUrl = domain + url + "?imageMogr2/format/" + this.Suffix + "/thumbnail/" + compress + "/quality/" + net + "!"
	default:
		imageUrl = domain + url
	}

	return imageUrl, nil
}
