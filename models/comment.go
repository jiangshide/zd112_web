package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type Comment struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId int64 `json:"contentId"`//动态ID
	Content string `json:"content"`//内容
	Urls string `json:"urls"`//内容Url/音频波文件数组地址
	Status int `json:"status"`//状态:1~正常评论,2~优秀评论,-1~异常评论,-4~关闭/折叠,-5~被投诉
	Reason string `json:"reason"`//原由
}

var COMMENT_STATUS=COMMENT+COMMENT_LEFT+"WHERE coalesce(C.status,0) IN(1,2) AND C.content_id=? "+COMMENT_ORDER_BY//评论列表:uid,uid,uid,content_id,limit,offset
var COMMENT_LEFT="FROM zd_comment C LEFT JOIN zd_user_profile UP ON UP.id=C.uid "		

var COMMENT = "SELECT C.id,C.uid,C.content_id contentId,C.content,C.urls,C.status,C.reason,C.create_time date,UP.nick,UP.icon,(SELECT name FROM zd_user_remarks WHERE uid=C.uid AND from_id=? ) remark,(SELECT count(1) FROM zd_comment_uid WHERE comment_id=C.id) num,(SELECT count(1) FROM zd_comment_praise WHERE comment_id=C.id AND coalesce(status,0) IN(1,2)) praiseNum,(SELECT status FROM zd_comment_praise WHERE comment_id=C.id AND uid=? ) praises,(SELECT reason FROM zd_report WHERE content_id=C.id AND source=4 AND uid=? ) reportr "//uid,uid,uid
var COMMENT_ORDER_BY="ORDER BY C.create_time DESC LIMIT ? OFFSET ? "

var COMMENT_ADD = "INSERT INTO zd_comment(uid,content_id,content,urls,status,reason)VALUES(?,?,?,?,?,?)"//添加评论

func (this *Comment) TableName() string {
	return TableName("comment")
}

func (this *Comment) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}


type CommentPraise struct {
	Id int64	`json:"id"`
	Uid int64	`json:"uid"` //用户ID
	CommentId int64 `json"commentId"` //评论ID
	Status int `json:"status"`//状态:-1~取消,1~点赞,2~超级赞
}

func (this *CommentPraise) TableName() string {
	return TableName("comment_praise")
}

func (this *CommentPraise) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *CommentPraise) ReadOrCreates()(created bool, id int64, err error){
	commentPraise:=CommentPraise{Id:this.Id,Uid:this.Uid,CommentId:this.CommentId}
	created,id,err = orm.NewOrm().ReadOrCreate(&commentPraise,"id","uid","comment_id")
	commentPraise.Status=this.Status
	id,err = commentPraise.Update()
	return
}

type CommentUid struct {
	Id int64	`json:"id"`
	CommentId int64 `json:"commentId"`//评论ID
	Uid int64 `json:"uid"`//用户ID
	FromUid int64 `json:"fromUid"`//来自用户ID
	ContentId int64 `json:"contentId"`//动态ID
	Content string `json:"content"`//内容
	Urls string `json:"urls"`//内容Url/音频波文件数组地址
	Status int `json:"status"`//状态:1~正常评论,-1~异常评论,2~优秀评论,-4~关闭/折叠,-5~被投诉
	Reason string `json:"reason"`//原由
}

var COMMENT_UID_STATUS=COMMENT_UID+COMMENT_UID_LEFT+"WHERE coalesce(CU.status,0) IN(1,2) AND CU.comment_id=? "+COMMENT_UID_ORDER_BY//评论列表:uid,uid,uid,comment_id,limit,offset
var COMMENT_UID_LEFT="FROM zd_comment_uid CU LEFT JOIN zd_user_profile UP ON UP.id=CU.uid "		

var COMMENT_UID = "SELECT CU.id,CU.comment_id,CU.uid,CU.from_uid fromUid,CU.content_id contentId,CU.content,CU.urls,CU.reply,CU.status,CU.reason,CU.create_time date,UP.nick,UP.icon,(SELECT coalesce((SELECT name FROM zd_user_remarks WHERE uid=CU.uid),nick) FROM zd_user_profile WHERE id=CU.uid) replyNick,(SELECT name FROM zd_user_remarks WHERE uid=CU.uid AND from_id=? ) remark,(SELECT count(1) FROM zd_comment_uid WHERE from_uid=CU.uid) num,(SELECT count(1) FROM zd_comment_uid_praise WHERE comment_uid_id=CU.id AND coalesce(status,0) IN(1,2)) praiseNum,(SELECT status FROM zd_comment_uid_praise WHERE comment_uid_id=CU.id AND uid=? ) praises,(SELECT reason FROM zd_report WHERE content_id=CU.id AND source=5 AND uid=? ) reportr "//uid,uid,uid
var COMMENT_UID_ORDER_BY="ORDER BY CU.create_time DESC LIMIT ? OFFSET ? "

var COMMENT_UID_ADD = "INSERT INTO zd_comment_uid(comment_id,uid,from_uid,content_id,content,urls,reply,status,reason)VALUES(?,?,?,?,?,?,?,?,?)"//针对用户添加评论


func (this *CommentUid) TableName() string {
	return TableName("comment_uid")
}

type CommentUidPraise struct {
	Id int64	`json:"id"`
	Uid int64	`json:"uid"` //用户ID
	CommentUidId int64 `json"commentUidId"` //评论ID
	Status int `json:"status"`//状态:-1~取消,1~点赞,2~超级赞
}

func (this *CommentUidPraise) TableName() string {
	return TableName("comment_uid_praise")
}

func (this *CommentUidPraise) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *CommentUidPraise) ReadOrCreates()(created bool, id int64, err error){
	commentUidPraise:=CommentUidPraise{Id:this.Id,Uid:this.Uid,CommentUidId:this.CommentUidId}
	created,id,err = orm.NewOrm().ReadOrCreate(&commentUidPraise,"id","uid","comment_uid_id")
	commentUidPraise.Status = this.Status
	id,err = commentUidPraise.Update()
	return
}

