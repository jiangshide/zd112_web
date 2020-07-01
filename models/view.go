package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type View struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId   int64  `json:"contentId"`//内容来源ID
	Num int `json:""num`//内容详情位置:来自哪张图片
	Type int `json:"type"`//内容类型:1~用户,2~频道,3~动态,4~评论
}

func (this *View) TableName() string {
	return TableName("view")
}

func (this *View) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

