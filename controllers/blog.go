package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sanskrit_web/models"
	// "strconv"
	"github.com/jiangshide/GoComm/utils"
)

type BlogController struct {
	BaseController
}

func (this *BlogController) Get(){
	format := this.getInt("format",-1)
	channelId := this.getInt("channelId",-1)
	jsd := this.getInt("jsd",-1)
	uid := this.userId

	page := this.page
	pageSize := this.pageSize

	if jsd > 0{
		page = 1
		pageSize=0
	}

	beego.Info("--------format:",format," | channelId:",channelId," | jsd:",jsd)
	if format > -1{
		blogs,ids,err := models.Blogs(this.userId,models.BLOG_FORMAT,[...]interface{}{uid,uid,uid,uid,uid,format,page,pageSize})
		beego.Info("--------blogs:",blogs," | ids:",ids," | err:",err)
		if err != nil || ids == 0{
			// this.false(-1,"data is null")	
		}else{
			for _,v := range *blogs{
				v["recommends"] =recommend(v["uid"])
				v["praiseNum"] = utils.ShowNum(v["praiseNum"].(string))
				v["viewNum"] = utils.ShowNum(v["viewNum"].(string))
				v["commentNum"] = utils.ShowNum(v["commentNum"].(string))
				v["shareNum"] = utils.ShowNum(v["shareNum"].(string))
				v["date"]= utils.StrTime(utils.TimeStamp(v["date"].(string)))
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

func (this *BlogController) AjaxFollow() {
	follow := new(models.UserFollow)
	follow.Uid = this.getInt64("uid",-1)
	follow.FromId = this.userId
	if status,_,err := follow.ReadOrCreates();err != nil{
		this.ajaxMsg(err,MSG_ERR)
	}else{
		this.ajaxMsg(status,MSG_OK)
	}
}

func (this *BlogController) AjaxPraise() {
	praise := new(models.BlogPraise)
	praise.Uid = this.userId
	praise.BlogId = this.getInt64("id",-1)
	if status,_,err := praise.ReadOrCreates();err != nil{
		this.ajaxMsg(err,MSG_ERR)
	}else{
		beego.Info("-------jsd~status:",status)
		this.ajaxMsg(status,MSG_OK)
	}
}


func (this *BlogController) AjaxCommendAdd(){
	contentId := this.getInt64("contentId",-1)
	content := this.getString("content","",0)
	urls := this.getString("urls","",0)
	status := this.getInt("status",1)
	reason := this.getString("reason","",0)
	if _,err := models.SqlRaw(models.COMMENT_ADD,[...]interface{}{this.userId,contentId,content,urls,status,reason});err != nil{
		this.ajaxMsg(err,MSG_ERR)
	}else{
		// total := new(Total)
		// comment:= make(map[string]interface{})
		// comment["content_id"]=contentId
		// comment["status"]=1
		// commentNum,_ := models.SqlCount("zd_comment",comment)

		// commentUid := make(map[string]interface{})
		// commentUid["content_id"]=contentId
		// commentUid["status"]=1//成功的状态
		// commentUidNum,_ := models.SqlCount("zd_comment_uid",commentUid)
		
		// total.Count = commentNum+commentUidNum
		// total.Size = commentUidNum
		this.ajaxMsg(status,MSG_OK)
	}
}









