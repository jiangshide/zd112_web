package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"	
 // "encoding/json"
 )

type Praise struct{
	Id int64	`json:"id"`
	Uid int64	`json:"uid"` //用户ID
	ContentId int64 `json"contentId"` //内容来源ID
	Status int `json:"status"`//状态:-1~取消,1~点赞,2~超级赞
	Type int `json:"type"`//内容类型:1~用户,2~频道,3~动态,4~评论
}

func (this *Praise) TableName()string{
	return TableName("praise")
}

func (this *Praise) ReadOrCreates(contentId,uid int64)(created bool, id int64, err error){
	praise:=Praise{ContentId:contentId,Uid:uid}
	created,id,err = orm.NewOrm().ReadOrCreate(&praise,"uid","content_id")
	return
}

func (this *Praise) Update()(int64,error){
	return orm.NewOrm().Update(this)
}

