package models

import (
	"github.com/astaxie/beego/orm"
)

type App struct {
	Id         int64 `json:"id"`
	Uid int64 `json:"uid"`
	Name       string `json:"name"` //应用名称
	Pkg string `json:"pkg"`//应用包名
	Platform   string   `json:"platform"` //平台
	Channel    string `json:"channel"`//渠道
	Version    string `json:"version"`//版本
	Code int `json:"code"`//版本code
	Env int `json:"env"`//环境:2~production,1~pre,0~qa
	Build int `json:"build"`//打包状态:1:打包中,2:打包成功,3:打包失败,4,无
	Cmd string `json:"cmd"`//打包脚本
	Log string `json:"log"`//打包日志详情
	Reason string `json:"reason"`//原由
	Duration int `json:"duration"`//提示安装时间间隔
	Times int `json:"times"`//提示安装时间次数
	Des   string `json:"des"` //更新内容提示
	Url     string   `json:"url"` //下载地址
	Qr string `json:"qr"`//二维码:可扫描下载
	Size int64 `json:"size"`//文件尺寸
	Status     int   `json:"status"` //更新状态:-2~删除-1~不可用,0~普通提示更新,1~提示强制更新,2～后台自动下载后更新(非静默更新),3~静默更新
	Count      int64  `json:"count"`//更新次数
}	

var APP_UPDATE = "INSERT INTO zd_app_update(uid,name,pkg,platform,channel,version,code,evn,duration,times,des,url,status)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)"//添加apk

var UPDATE_NAME = SqlUpdate+"WHERE instr(name,?) >0 order by update_time DESC limit ? offset ?"

var UPDATE_APK = SqlUpdate+"WHERE name=? and platform=? androidnd channel=? and version=? "

var SqlUpdate = "SELECT id,uid,name,platform,channel,version,code,duration,times,des,url,status,count,update_time date FROM zd_app_update "

var UPDATE_STATUS = "UPDATE zd_app_update SET status=? WHERE id=?"
var UPDATE_DEL = "DELETE FROM zd_app_update WHERE id=?"



func (this *App) TableName() string {
	return TableName("app")
}

func (this *App) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *App) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *App) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *App) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this.TableName()).OrderBy("-id")
}

func (this *App) List(pageSize, offSet int) (list []*App, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type AppCount struct{
	Id int64 `json:"id"`//
	AppId int64 `json:"appId"`//应用ID
	Status int `json:"status"`//更新状态:0~请求更新,2~下载更新文件,3~安装文件,4~安装成功,-1~请求更新失败,-2~下载更新文件失败,-3~安装文件失败
}

func (this *AppCount) TableName() string {
	return TableName("app_count")
}

