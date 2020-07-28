package main

import (
	_ "sanskrit_web/routers"
	"github.com/astaxie/beego"
	"sanskrit_web/models"
)

func main() {
	models.Init()
	beego.Run()
}

