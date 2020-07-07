package controllers

import (
	"zd112_web/models"
	"github.com/astaxie/beego"
)

type TestController struct{
	beego.Controller
}

func(this *TestController)Get(){
	this.TplName = "test.html"
}

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.channel()
	this.display("index")
}

func (this *IndexController) channel(){
	id := this.getInt("id",0)
	uid := this.userId
	beego.Info("---------jsd~maps~id:",id," | uid:",uid)
	if id == 1{
		maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_UID,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
		beego.Info("maps---->:",maps," | id:",id," | err:",err)
		if err != nil || id==0{
			maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_UNFOLLOW,[...]interface{}{uid,uid,uid,uid,this.page,this.pageSize})
			beego.Info("maps---->:",maps," | id:",id," | err:",err)
			if err != nil || id==0{
				maps,id,err := models.Channels(uid,models.CHANNEL_RECOMMEND_ALL,[...]interface{}{uid,uid,uid,uid,uid,this.page,this.pageSize})
				beego.Info("maps---->:",maps," | id:",id," | err:",err)
				if err != nil || id == 0{
					// this.showTips(err)
				}else{
					this.Data["maps"] = maps
				}
			}else{
				this.Data["maps"] = maps	
			}
		}else{
			this.Data["maps"]=maps
		}
	}else if id == 2{
		maps,id,err := models.Channels(uid,models.CHANNEL_HOT,[...]interface{}{uid,uid,uid,uid,this.page,this.pageSize})
		beego.Info("maps---->:",maps," | id:",id," | err:",err)
		if err != nil || id==0{
			// this.showTips("data is null")
		}else{
			this.Data["maps"]=maps	
		}
	}else{
		maps,id,err := models.Channels(uid,models.CHANNEL_TYPE_ID,[...]interface{}{uid,uid,uid,uid,id,this.page,this.pageSize})
		beego.Info("maps---->:",maps," | id:",id," | err:",err)
		if err != nil || id==0{
			// this.showTips("data is null")
		}else{
			this.Data["maps"]=maps
		}
	}
}
