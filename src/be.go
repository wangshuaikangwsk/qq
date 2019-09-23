package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello 谷亚")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}