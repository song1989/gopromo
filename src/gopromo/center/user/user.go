package userCenter

import (
	"encoding/json"
	"log"

	"gopromo/center"
	"gopromo/center/thrift/userinfoservice"
)

const (
	host = "10.10.106.45"
	port = "9082"
)

func New() userCententer {
	transport, protocolFactory := center.GetFactory(host, port)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", host+":"+port)
	}
	//defer transport.Close()

	client := userinfoservice.NewUserInfoServiceClientFactory(transport, protocolFactory)
	return &userCentent{client}
}

type userCententer interface {
	GetUserInfoById(id int64) (userField, error)
	GetUserInfoByPrivateKey(privateKey string) (userField, error)
}

type userCentent struct {
	Client *userinfoservice.UserInfoServiceClient
}

type thriftData struct {
	Succ bool   `json:"succ"`
	Code string `json:"code"`
	Data userField
}

type userField struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Number     string `json:"number"`
	PrivateKey string `json:"-"` //omitempty 如果为空置则忽略字段
}

func (this *userCentent) GetUserInfoById(id int64) (userField, error) {
	d, err := this.Client.GetUserInfoById(id)
	//defer this.Client.Transport.Close()
	if err != nil {
		log.Println("userinfoservice:", err)
	}
	log.Println(d)

	var thriftValue thriftData
	if err := json.Unmarshal([]byte(d), &thriftValue); err != nil {
		return userField{}, err
	}
	if thriftValue.Succ == false {
		return userField{}, nil
	}
	return thriftValue.Data, nil
}

func (this *userCentent) GetUserInfoByPrivateKey(privateKey string) (userField, error) {
	d, err := this.Client.GetUserInfoByPrivateKey(privateKey)
	//defer this.Client.Transport.Close()
	if err != nil {
		log.Println("userinfoservice:", err)
	}
	var thriftValue thriftData
	if err := json.Unmarshal([]byte(d), &thriftValue); err != nil {
		return userField{}, err
	}
	if thriftValue.Succ == false {
		return userField{}, nil
	}
	return thriftValue.Data, nil
}
