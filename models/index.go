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

type ChannelType struct {
	// Id         int64
	// Name       string
	// Action     string
	// Descript   string
	// CreateId   int64
	// UpdateId   int64
	// CreateTime int64
	// UpdateTime int64
	// Views      int64
	Id int `json:"id"`
	Name string `json:"name"`//频道类型名称
	Des string `json:"des"`//频道类型描述
}

func (this *ChannelType) TableName() string {
	return TableName("channel_type")
}

func (this *ChannelType) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *ChannelType) List(pageSize, offSet int) (list [] *ChannelType, total int64) {
	// query := orm.NewOrm().QueryTable(this.TableName())
	// total, _ = query.Count()
	// query.Limit(pageSize, offSet).All(&list)
	return
}
