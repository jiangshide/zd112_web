package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type Top struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId   int64  `json:"contentId"`//内容来源ID
	Status int `json:"status"`//状态:1~置顶,0~取消
	Especially int 	`json:"especially"`//特别状态:1~超级顶,0～取消超级顶
	Type int `json:"type"`//内容类型:1~用户,2~频道,3~动态,4~评论
}

func (this *Top) TableName() string {
	return TableName("top")
}

func (this *Top) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

