package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"	
 )

func InitFormat(){
	var list = [...]string{"图文","语音","视频","文字","Web","VR"}
	for _,v := range list{
		format := new(Format)
		format.Name = v
		format.Des = v
		format.Uid = 0
		id,err := format.Add()
		beego.Info("id:",id," | err:",err)
	}
}

type Format struct {
	Id int64	`json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	Name string `json:"name"`//内容类型名称
	Des string `json:"des"`//内容类型描述	
}

func (this *Format) TableName() string {
	return TableName("format")
}

func (this *Format) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

var FORMAT_UPDATE = "UPDATE zd_format SET des=? WHERE id=?"

func SqlFormat(where string)string{
	// return "SELECT id,name,des,uid,update_time as date FROM zd_format "+where+" order by update_time DESC limit ? offset ?"
	return "SELECT id,name FROM zd_format "+where+" order by update_time DESC limit ? offset ?"
}

