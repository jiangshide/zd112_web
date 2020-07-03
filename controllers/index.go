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
	this.nav()
	this.banner()
	this.qrImg()
	this.group()
	this.universityItem()
	this.channel()
	this.nationItem()
	this.audioItem()
	this.videoItem()
	this.originalItem()
	this.questionItem()
	this.display("index")
}

func (this *IndexController) channel(){
	id := this.getInt("id",-1)
	uid := this.userId
	maps,ids,err := models.Channels(uid,models.CHANNEL_TYPE_ID,[...]interface{}{uid,uid,uid,uid,id,20,0})
	beego.Info("------maps:",maps," | id:",id," | uid:",uid)
	if err != nil || ids==0{
		// this.showTips("data is null")
	}else{
		this.Data["maps"] = maps
	}
}

func (this *IndexController) group() {
	group := map[string]string{"/static/img/group3.png": "最新","/static/img/group1.png": "最热", "/static/img/group2.png": "猜你喜欢", "/static/img/group4.png": "每日签到"}
	this.Data["group"] = group
}

func (this *IndexController) universityItem() {
	this.Data["university"] = this.university(0)
	university := map[string]string{"/static/img/item1.png": "一剪梅", "/static/img/item2.png": "古来黎", "/static/img/item3.png": "李白", "/static/img/item4.png": "崧台"}
	this.Data["universityItem"] = this.getItem(university)
}

func (this *IndexController) nationItem() {
	this.Data["nation"] = this.nation(0, 9)
	nation := map[string]string{"/static/img/nation1.png": "侗族", "/static/img/nation2.png": "土家族", "/static/img/nation3.png": "羌族", "/static/img/nation4.png": "汉族"}
	this.Data["nationItem"] = this.getItem(nation)
}

func (this *IndexController) audioItem() {
	this.Data["audio"] = this.audio(0)
	audio := map[string]string{"/static/img/music1.png": "李白", "/static/img/music2.png": "陶渊明", "/static/img/music3.png": "滔滔", "/static/img/music4.png": "红海"}
	this.Data["audioItem"] = this.getItem(audio)
}

func (this *IndexController) videoItem() {
	this.Data["video"] = this.video(0)
	video := map[string]string{"/static/img/video1.png": "告别", "/static/img/video2.png": "美丽的春天", "/static/img/video3.png": "好日子", "/static/img/video4.png": "未来"}
	this.Data["videoItem"] = this.getItem(video)
}

func (this *IndexController) originalItem() {
	this.Data["original"] = this.original(0)
	original := map[string]string{"/static/img/original1.png": "黎明", "/static/img/original2.png": "科技", "/static/img/original3.png": "怎么学习", "/static/img/original4.png": "地球原理"}
	this.Data["originalItem"] = this.getItem(original)
}

func (this *IndexController) questionItem() {
	this.Data["question"] = this.question(0)
	question := map[string]string{"/static/img/nation3.png": "牡蛎", "/static/img/music2.png": "详探", "/static/img/video4.png": "高数", "/static/img/original1.png": "机器学"}
	this.Data["questionItem"] = this.getItem(question)
}

func (this *IndexController) getItem(poetry map[string]string) []*Content {
	contentArr := make([] *Content, 0)
	for k, v := range poetry {
		content := new(Content)
		content.Id = 1
		content.Title = v
		content.Name = "李清照"
		content.Year = "宋代"
		content.Icon = k
		content.Like = "20万"
		content.Follow = "100万"
		content.Info = 80
		content.Before = "1天前"
		user := new(User)
		user.Id = 1
		user.Name = "春天来到"
		user.Icon = "/static/img/user_icon.png"
		content.User = user
		contentArr = append(contentArr, content)
	}
	return contentArr
}	

func (this *IndexController) banner() {
	// banner := new(models.Banner)
	// list, _ := banner.List(this.pageSize, this.offSet)
	// maps ,_,_ := models.SqlList("SELECT * FROM zd_channel WHERE format=0",[...]interface{}{})
	// this.Data["banner"] = maps

	uid := this.userId
	maps,id,err := models.Channels(uid,models.CHANNEL_OFFICIAL,[...]interface{}{uid,uid,uid,uid,20,0})
	beego.Info("maps:",maps)
	if err != nil || id==0{
		this.showTips("data is null")		
	}else{
		this.Data["banner"] = maps
	}
}

func (this *IndexController) qrImg() {
	app := new(models.App)
	res,_ := app.List(1,0)
	beego.Info("app:",app)
	this.Data["android"] = res[0].Qr
}
