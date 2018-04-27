package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {

	db, err := gorm.Open("mysqlConnect", "root:root123@/in_promo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//创建
	//	user := User{
	//		Name: "test2",
	//		Age:  20,
	//	}
	//	db.Create(&user)
	//	fmt.Println("insrt ok")

	//查询
	var users []User
	result := db.Find(&users)
	if err := result.Error; err != nil {
		fmt.Println("该数据不存在")
	}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Name)
	}

	//	user := User{}
	//	data := db.Find(&user)
	//fmt.Println(data)
}
