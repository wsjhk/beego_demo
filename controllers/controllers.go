package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "www.baidu.com"
	this.Data["Email"] = "admin@admin.com"
	this.TplName = "index.tpl"
	// this.Ctx.WriteString("hello")
}
