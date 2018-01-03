package main

import (
	_ "zd112_web/routers"
	"github.com/astaxie/beego"
	"zd112_web/models"
)

func main() {
	models.Init()
	beego.Run()
}

