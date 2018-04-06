package models

import "github.com/astaxie/beego/orm"

type Banner struct {
	Id         int64
	Name       string
	Link       string
	Icon       string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int
}

func (this *Banner) TableName() string {
	return TableName("web_banner")
}

func (this *Banner) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Banner) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Banner) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Banner) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Banner) List(pageSize, offSet int) (list []*Banner, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type Nav struct {
	Id         int64
	Name       string
	Action     string
	Descript   string
	CreateId   int64
	UpdateId   int64
	CreateTime int64
	UpdateTime int64
	Views      int64
}

func (this *Nav) TableName() string {
	return TableName("web_nav")
}

func (this *Nav) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Nav) List(pageSize, offSet int) (list [] *Nav, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.Limit(pageSize, offSet).All(&list)
	return
}
