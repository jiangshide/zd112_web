package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"	
 // "encoding/json"
 )

type ReportType struct{
	Id int `json:"id"`
	Name string `json:"name"`//投诉类型名称
}

var REPORT_TYPE_ADD = "INSERT INTO zd_report_type(name)VALUES(?)"//投诉类型

func InitReportType(){
	var list = [...]string{"谩骂造谣","广告传销","内容抄袭","诈骗","色情","暴力","版权盗取","反动","违法信息"}
	for _,v := range list{
		_,err := SqlRaw(REPORT_TYPE_ADD,[...]interface{}{v})
		if err != nil{
			beego.Info(err)
		}		
	}
}

var REPORT_TYPE = "SELECT id,name FROM zd_report_type"

type Report struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID	
	ContentId   int64  `json:"contentId"`//内容来源ID
	Type int `json:"type"`//投诉类型:
	Source int `json:"source"`//内容来源类型:1~用户,2~频道,3~动态,4~评论,5~用户评论的评论
	Status int `json:"status"`//状态:0~未处理,1～成功,-1~失败
	Reason string 	`json:"reason"`//原由
}

func (this *Report) TableName() string {		
	return TableName("report")
}

func (this *Report) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Report) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Report) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Report) Query() error {
	return orm.NewOrm().QueryTable(this.TableName()).Filter("uid",this.Uid).Filter("content_id",this.ContentId).Filter("source",this.Source).One(this)
}

func (this *Report) ReadOrCreates(uid,contentId int64,types,source,status int,reason string)(created bool, id int64, err error){
	report:=Report{Uid:uid,ContentId:contentId}
	created,id,err = orm.NewOrm().ReadOrCreate(&report,"uid","content_id")
	if !created {
		report.Type = types
		report.Source = source
		report.Status=status
		report.Reason=reason
		id,err = report.Update()
	}
	return
}

var REPORT_CONTENTID="ADD content_id=?"
var REPORT_UID="ADD uid=? "
var REPORT_STATUS="WHERE status=? "
var SqlReport="SELECT id,uid,content_id contentId,type,source,status,reason FROM zd_report "

var REPORT_ADD = "REPLACE INTO zd_report(id,uid,content_id,type,source,status,reason)VALUES(?,?,?,?,?,?,?)"//举报


