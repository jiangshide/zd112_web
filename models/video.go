package models

import "github.com/astaxie/beego/orm"

type Video struct {
	Id         int64
	Name       string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
}

func (this *Video) TableName() string {
	return TableName("web_video")
}

func (this *Video) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Video) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Video) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Video) Query() error {
	if this.Id != 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}
