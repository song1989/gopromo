package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func (this *MainController) Post() {
	this.Ctx.WriteString("hello world Post")
}

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("get hello")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/hello", &HelloController{})
	beego.Run()
}
