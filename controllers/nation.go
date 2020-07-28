package controllers

import (
	"sanskrit_web/models"
	"github.com/astaxie/beego"
)

type NationController struct{
	BaseController
}

func (this *NationController) Get(){
	this.display("nation/nation")
}

func (this *NationController) Names(){
	nation := new(models.Nation)
	result,count := nation.List(this.pageSize,this.offSet)
	beego.Info("result:",result," | count:",count)
}