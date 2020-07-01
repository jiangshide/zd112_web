package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type Share struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId   int64  `json:"contentId"`//内容来源ID
	Status int `json:"status"`//状态:1~成功,0～失败
	Type int `json:"type"`//内容类型:1~用户,2~频道,3~动态,4~评论
	FromId int64 `json:"fromId"`//来自用户ID
}

func (this *Share) TableName() string {
	return TableName("share")
}

func (this *Share) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

