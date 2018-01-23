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

	this.auth()
}

func (this *BaseController) setCook(user *models.User, time int) {
	this.Ctx.SetCookie(AUTH, fmt.Sprint(user.Id)+"|"+utils.Md5(this.getIp()+"|"+user.Password+user.Salt), time)
}

func (this *BaseController) auth() (user *models.User, err error) {
	cook := this.Ctx.GetCookie(AUTH)
	beego.Info("-----cook:", cook)
	if strings.Contains(cook, "|") {
		cookArr := strings.Split(cook, "|")
		user = new(models.User)
		user.Id, _ = strconv.ParseInt(cookArr[0], 11, 64)
		if err = user.Query(); err == nil {
			this.userId = user.Id
			this.userName = user.Name
			this.Data["userId"] = this.userId
			this.Data["userName"] = this.userName
			this.Data["isLogin"] = true
			beego.Info("------userId:", this.userId, " | userName:", this.userName)
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

/**
comm data
 */
type Content struct {
	Id     int
	Title  string
	Name   string
	Year   string
	Icon   string
	Like   string
	Follow string
	Info   int
	Before string
	User   *User
}

type User struct {
	Id   int
	Name string
	Icon string
}

func (this *BaseController) list(arr ...string) []map[int]interface{} {
	nation := make([]map[int]interface{}, 0)
	na := make(map[int]interface{}, 0)
	for k, v := range arr {
		na[k] = v
		t := k + 1
		if t%10 == 0 {
			nation = append(nation, na)
			na = nil
			na = make(map[int]interface{}, 0)
		}
		if k+1 == len(arr) {
			nation = append(nation, na)
		}
	}
	return nation
}

func (this *BaseController) university(index int) []string {
	university := []string{"学府", "诗词", "史书典籍", "诗词周边", "美文阅读", "书法欣赏", "其它"}
	return university[index:]
}

func (this *BaseController) nation(index, end int) []string {
	nation := []string{"民族", "阿昌族", "白族", "保安族", "布朗族", "布依族", "藏族", "朝鲜族", "达翰尔族", "傣族", "昂德族", "东乡族", "侗族", "独龙族", "俄罗斯族", "鄂伦春族", "鄂温克族", "高山族", "哈尼族", "哈萨克族", "汉族", "赫哲族", "回族", "基诺族", "京族",
		"景颇族", "柯尔克孜族", "拉祜族", "黎族", "傈僳族", "珞巴族", "满族", "毛南族", "门巴族", "蒙古族", "苗族", "仫佬族", "纳西族", "怒族", "普米族", "羌族", "撒拉族", "畲族", "水族", "塔吉克族", "塔塔尔族", "土家族", "图族", "佤族", "维吾尔族", "乌孜别克族", "锡伯族", "瑶族", "彝族", "仡佬族", "裕固族", "壮族"}
	if end > 0 {
		return nation[index:end]
	}
	return nation[index:]
}

func (this *BaseController) audio(index int) []string {
	audio := []string{"音乐", "古琴", "琵琶", "古筝", "笛子", "葫芦丝", "芦笋", "现代流行曲", "古典", "经典"}
	return audio[index:]
}

func (this *BaseController) video(index int) []string {
	video := []string{"视频", "短视频", "长视频", "MV"}
	return video[index:]
}

func (this *BaseController) original(index int) []string {
	original := []string{"原创", "手工艺", "摄影", "纯艺术", "服装", "视频", "音乐"}
	return original[index:]
}

func (this *BaseController) question(index int) []string {
	question := []string{"题库", "英语", "数学", "物理", "化学", "政治", "生物", "地理", "语文", "历史"}
	return question[index:]
}

type Nav struct {
	Name    string
	Action  string
	NavList [] NavList
}

type NavList struct {
	Name        string
	Action      string
	NavListItem [] NavListItem
}

type NavListItem struct {
	Name   string
	Action string
}

func (this *BaseController) nav() {
	navstr := []string{"首页", "学府", "民族", "音乐", "视频", "原创", "题库"}
	navArr := []Nav{}
	for _, v := range navstr {
		var nav Nav
		nav.Name = v
		nav.Action = "/university"
		if v == "学府" {
			for _, v := range this.university(1) {
				var navList NavList
				navList.Name = v
				if v == "诗词" {
					shichi := []string{"年代诗人", "著名诗人", "诗词标签", "诗词故事"}
					for _, v := range shichi {
						var navListItem NavListItem
						navListItem.Name = v
						navList.NavListItem = append(navList.NavListItem, navListItem)
					}
				} else {
					shichi := []string{"年代", "著名", "诗词标签", "诗词故事"}
					for _, v := range shichi {
						var navListItem NavListItem
						navListItem.Name = v
						navList.NavListItem = append(navList.NavListItem, navListItem)
					}
				}
				nav.NavList = append(nav.NavList, navList)
			}
		} else if v == "民族" {
			for _, v := range this.nation(1, 0) {
				var navList NavList
				navList.Name = v
				nav.NavList = append(nav.NavList, navList)
			}
		} else if v == "音乐" {
			for _, v := range this.audio(1) {
				var navList NavList
				navList.Name = v
				nav.NavList = append(nav.NavList, navList)
			}
		} else if v == "视频" {
			for _, v := range this.video(1) {
				var navList NavList
				navList.Name = v
				nav.NavList = append(nav.NavList, navList)
			}
		} else if v == "原创" {
			for _, v := range this.original(1) {
				var navList NavList
				navList.Name = v
				nav.NavList = append(nav.NavList, navList)
			}
		} else if v == "题库" {
			for _, v := range this.question(1) {
				var navList NavList
				navList.Name = v
				nav.NavList = append(nav.NavList, navList)
			}
		}
		navArr = append(navArr, nav)
	}
	this.Data["nav"] = navArr

	this.Data["navUniversity"] = this.university(1)
	this.Data["navNation"] = this.list(this.nation(1, 0)...)
	this.Data["navAudio"] = this.list(this.audio(1)...)
	this.Data["navVideo"] = this.video(1)
	this.Data["navOriginal"] = this.original(1)
	this.Data["navQuestion"] = this.question(1)
}
