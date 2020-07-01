package models

import ("github.com/astaxie/beego/orm"
//  "github.com/astaxie/beego"	
 )


type Word struct {
	Id int64	`json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	ContentId int64 `json:contentId`//内容来源ID
	Name string `json:"name"`//关键字名称
	Source int `json:"source"`//内容类型:1~用户,2~频道,3~动态,4~评论	
}

func (this *Word) TableName() string {		
	return TableName("word")
}

func (this *Word) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Word) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Word) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Word) ReadOrCreates(uid,contentId int64,name string,source int)(created bool, id int64, err error){
	word:=Word{Uid:uid,ContentId:contentId,Name:name,Source:source}
	created,id,err = orm.NewOrm().ReadOrCreate(&word,"uid","content_id","name","source")
	id,err = word.Update()
	return
}

var WORD_QUERY = "SELECT name,source FROM zd_word WHERE uid=? order by update_time DESC limit ? offset ?"

