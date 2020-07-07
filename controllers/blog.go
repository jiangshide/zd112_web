package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"zd112_web/models"
	// "strconv"
)

type BlogController struct {
	BaseController
}

func (this *BlogController) Get(){
	format := this.getInt("format",-1)
	channelId := this.getInt("channelId",-1)
	uid := this.userId

	beego.Info("--------format:",format," | channelId:",channelId)
	if format > -1{
		blogs,ids,err := models.Blogs(this.userId,models.BLOG_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,format,this.page,this.pageSize})
		beego.Info("--------blogs:",blogs," | ids:",ids," | err:",err)
		if err != nil || ids == 0{
			// this.false(-1,"data is null")	
		}else{
			for _,v := range *blogs{
				v["recommends"] =recommend(v["uid"])
				v["praiseNum"] = ShowNum(v["praiseNum"].(string))
				v["viewNum"] = ShowNum(v["viewNum"].(string))
				v["commentNum"] = ShowNum(v["commentNum"].(string))
				v["shareNum"] = ShowNum(v["shareNum"].(string))
				v["date"]= StrTime(TimeStamp(v["date"].(string)))
			}
			this.Data["blogs"]=blogs
		}
	}else if channelId > -1{
		blogs,ids,err := models.Blogs(this.userId,models.BLOG_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,format,this.page,this.pageSize})
		beego.Info("--------blogs:",blogs," | ids:",ids," | err:",err)
		if err != nil || ids == 0{
			// this.false(-1,"data is null")	
		}else{
			this.Data["blogs"]=blogs
		}
	}
	this.display("blog/blog")
}

func recommend(uid interface{})(list *[]orm.Params){
	maps,_,_ := models.SqlList("SELECT id,(SELECT name FROM zd_channel WHERE id=B.channel_id) channel,B.title,B.des,B.cover,B.name,B.format FROM zd_blog B WHERE uid = ? ORDER BY update_time DESC LIMIT ? OFFSET ? ",[...]interface{}{uid,6,0})
	return maps
}