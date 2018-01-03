package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/gwpp/tinify-go/tinify"
	"zd112_web/models"
	"fmt"
	"github.com/jiangshide/GoComm/utils"
	"strconv"
	"time"
	"net"
)

const (
	MSG_OK           = 0
	MSG_ERR          = -1
	AUTH             = "auth"
	DEFAULT_TIPS     = "该项不能为空!"
	DEFAULT_MIN_SIZE = 3
)

type BaseController struct {
	beego.Controller
	version    string
	controller string
	action     string
	userId     int64
	userName   string
	logo       string
	page       int
	pageSize   int
	offSet     int
	upload     string
}

func (this *BaseController) Prepare() {
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	beego.Info("controller:", this.controller, " | action:", this.action)
	this.Data["siteName"] = beego.AppConfig.String("site.app_name")
	this.version = beego.AppConfig.String("version")
	this.page, _ = this.GetInt("page", 1)
	this.pageSize, _ = this.GetInt("limit", 30)
	this.offSet = (this.page - 1) * this.pageSize
	Tinify.SetKey(this.GetString("pic_key"))
	this.upload = this.GetString("upload")
}

func (this *BaseController) setCook(user *models.User, time int) {
	this.Ctx.SetCookie(AUTH, fmt.Sprint(user.Id)+"|"+utils.Md5(this.getIp()+"|"+user.Password+user.Salt), time)
}

func (this *BaseController) auth() (user *models.User, err error) {
	cook := this.Ctx.GetCookie(AUTH)
	if strings.Contains(cook, "|") {
		cookArr := strings.Split(cook, "|")
		user = new(models.User)
		user.Id, _ = strconv.ParseInt(cookArr[0], 11, 64)
		if err = user.Query(); err == nil {
			this.userId = user.Id
			this.userName = user.Name
		}
	}
	return
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getIp() string {
	return this.Ctx.Input.IP()
}

func (this *BaseController) getMac() (mac string) {
	inter, err := net.InterfaceByName("eth0")
	if err == nil {
		mac = inter.HardwareAddr.String()
	}
	return
}

func (this *BaseController) getInt(key string, defaultValue int) int {
	resInt, err := this.GetInt(key, defaultValue)
	if err != nil {
		this.ajaxMsg(err, MSG_ERR)
	}
	return resInt
}

func (this *BaseController) getInt64(key string, defaultValue int64) int64 {
	resInt, err := this.GetInt64(key, defaultValue)
	if err != nil {
		this.ajaxMsg(err, MSG_ERR)
	}
	return resInt
}

func (this *BaseController) getId(defaultValue int) int {
	return this.getInt("id", defaultValue)
}

func (this *BaseController) getId64(defaultValue int64) int64 {
	return this.getInt64("id", defaultValue)
}

func (this *BaseController) getGroupId(defaultvalue int) int {
	return this.getInt("groupId", defaultvalue)
}

func (this *BaseController) getGroupId64(defaultvalue int64) int64 {
	return this.getInt64("groupId", defaultvalue)
}

func (this *BaseController) getString(key, tips string, minSize int) (value string) {
	value = strings.TrimSpace(this.GetString(key, ""))
	errorMsg := ""
	if len(value) == 0 {
		errorMsg = tips
	} else if len(value) < minSize {
		errorMsg = "长度不能小于" + fmt.Sprint(minSize) + "字符:" + value
	}
	if errorMsg != "" {
		this.ajaxMsg(errorMsg, MSG_ERR)
	}
	return
}

func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) display(tpl ...string) {
	var tplName string
	if len(tpl) > 0 {
		tplName = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplName = this.controller + "/" + this.action + ".html"
	}
	beego.Info("------tpl:", tpl, " | tplName:", tplName)
	this.TplName = tplName
}

func (this *BaseController) ajaxMsg(msg interface{}, code int) {
	out := make(map[string]interface{})
	out["status"] = code
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxList(msg interface{}, code int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.Data["data"] = data
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ajaxMsgFile(msg interface{}, size, reSize int64, code int) {
	out := make(map[string]interface{})
	out["status"] = code
	out["message"] = msg
	out["size"] = size
	out["resize"] = reSize
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) showTips(errorMsg interface{}) {
	flash := beego.NewFlash()
	flash.Error(fmt.Sprint(errorMsg))
	flash.Store(&this.Controller)
	controller, action := this.GetControllerAndAction()
	beego.Info("--------controller:", controller, " | action:", action)
	this.Redirect(beego.URLFor(controller+"."+action), 302)
	this.StopRun()
}

func (this *BaseController) Upload() {
	f, fh, err := this.GetFile("file")
	defer f.Close()
	fileName := fh.Filename
	sufix := "default"
	if strings.Contains(fh.Filename, ".") {
		sufix = fileName[strings.LastIndex(fileName, ".")+1:]
	}
	fileName = utils.Md5(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	toFilePath := this.upload + sufix + "/" + fileName
	var size, reSize int64
	if err = this.SaveToFile("file", utils.GetCurrentDir(toFilePath)); err == nil {
		size, reSize = this.compress(toFilePath)
	}
	this.ajaxMsgFile(toFilePath, size, reSize, MSG_OK)
}

func (this *BaseController) setFileSize(row map[string]interface{}, file string) {
	size, _ := utils.FileSize(file)
	row["Size"] = size
}

func (this *BaseController) compress(path string) (int64, int64) {
	path = utils.GetCurrentDir(path)
	size, _ := utils.FileSize(path)
	src, err := Tinify.FromFile(path)
	var reSize int64
	if err == nil {
		if err = src.ToFile(path); err == nil {
			res, _ := utils.FileSize(path)
			reSize = res
		}
	}
	if err != nil {
		beego.Error("compress:", err)
	}
	return size, reSize
}
