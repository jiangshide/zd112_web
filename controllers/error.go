package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (this *ErrorController)Error404(){
	this.Data["content"]="正在开发中..."
	this.TplName=""
}