package models

import ("github.com/astaxie/beego/orm"
 // "github.com/astaxie/beego"
 // "encoding/json"
 )

type History struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ContentId   int64  `json:"blogId"`//内容来源ID
	Num int `json:"status"`//内容详情位置:来自哪张图片
	Source int `json:"source"`//内容类型:1~用户,2~频道,3~动态,4~评论,5~音乐
}

func (this *History) TableName() string {		
	return TableName("history")
}

func (this *History) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *History) ReadOrCreates()(created bool, id int64, err error){
	history := History{Uid:this.Uid,ContentId:this.ContentId}
	created,id,err = orm.NewOrm().ReadOrCreate(&history,"uid","content_id")
	history.Num=this.Num
	history.Source = this.Source
	id,err = history.Update()
	return
}