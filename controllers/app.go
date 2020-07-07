package controllers

import (
	"zd112_web/models"
	// "github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
	"github.com/jiangshide/GoComm/utils"
	"fmt"	
	// "strconv"
	// "time"
	"strings"
	// "runtime"
)

type AppController struct {
	BaseController
}

func (this *AppController) Add() {
	// 	ip := this.getIp()
	// 		mac := this.getMac()
	// 		device := runtime.GOOS
	// 		arch := runtime.GOARCH
	// 		sdkVersion := runtime.Version()
	// beego.Info("ip:",ip," | mac:",mac," | device:",device," | arch:",arch," | sdkVersion:",sdkVersion)
	if this.isPost() {
		app := new(models.App)
		url := this.getString("url","",0)
		app.Name = this.getString("name","",0)
		app.Pkg = "com.android.sanskrit"
		app.Channel = this.getString("channel","",0)
		app.Platform = this.getString("platform","",0)
		app.Env = this.getInt("env",0)
		app.Version = this.getString("version","",0)
		// app.Url = "http://" + utils.GetLocalAdder() + ":"+beego.AppConfig.String("httpport")+url
		app.Url = this.host+url
		app.Duration = this.getInt("internel",1)
		app.Times = this.getInt("times",10)
		app.Status = this.getInt("status",0)
		app.Des = this.getString("content","",0)
		imgUrl := url[:strings.LastIndex(url, ".")]+".png"
		app.Qr = imgUrl
		qrcode.WriteFile(app.Url, qrcode.Medium, 256, utils.GetCurrentDir(imgUrl))
		if _,err := app.Add();err != nil{
			this.showTips("添加失败:"+fmt.Sprint("%s",err))
		}else{
			this.redirect("/")
		}
	}else{
		this.display("app/app_add")
	}
}
	