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
	"math"
	"bytes"
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
	defaultIcon string
	host string
}

func (this *BaseController) Prepare() {
	controller, action := this.GetControllerAndAction()
	this.controller = strings.ToLower(controller[0:len(controller)-10])
	this.action = strings.ToLower(action)
	beego.Info("controller:", this.controller, " | action:", this.action)
	this.Data["siteName"] = beego.AppConfig.String("site.app_name")
	this.version = beego.AppConfig.String("version")
	this.page, _ = beego.AppConfig.Int("page")
	this.pageSize, _ = beego.AppConfig.Int("pageSize")
	key := beego.AppConfig.String("pic_key")
	Tinify.SetKey(key)
	this.upload = beego.AppConfig.String("upload")
	this.defaultIcon = "http://" + utils.GetLocalAdder() + ":"+beego.AppConfig.String("httpport")+"/static/img/logo.png"
	beego.Info("------upload:",this.upload," | key:",key," | page:",this.page)	
	this.host = "http://" + utils.GetLocalAdder() + ":"+beego.AppConfig.String("httpport")	
	this.Data["host"]=this.host
	this.Data["PIC"]=0
	this.Data["AUDIO"]=1
	this.Data["VIDEO"]=2
	this.page,this.pageSize = this.Limit(this.page,this.pageSize)
	this.auth()
	this.nav()
	this.banner()
	this.qrImg()
	this.group()
}

func (this *BaseController) Limit(page,pageSize int) (int, int) {
	num, _ := this.GetInt("page", page)
	size, _ := this.GetInt("pageSize", pageSize)
	return size, size*num
}

func (this *BaseController) setCook(user *models.User, time int) {
	this.Ctx.SetCookie(AUTH, fmt.Sprint(user.Id)+"|"+utils.Md5ToStr(this.getIp()+"|"+user.Password+user.Salt), time)
}

func (this *BaseController) auth() (profile *models.Profile, err error) {
	cook := this.Ctx.GetCookie(AUTH)
	beego.Info("-----cook:", cook)
	if strings.Contains(cook, "|") {
		cookArr := strings.Split(cook, "|")
		profile = new(models.Profile)
		id, _ := strconv.ParseInt(cookArr[0], 11, 64)
		if err = profile.Query("id",id); err == nil {
			this.userId = profile.Id
			this.userName = profile.Nick
			this.defaultIcon = profile.Icon
			this.Data["userId"] = this.userId
			this.Data["userName"] = this.userName
			this.Data["icon"]=this.defaultIcon
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
	fileName = utils.Md5ToStr(this.userName+time.RubyDate+utils.GetRandomString(10)) + "_" + fmt.Sprint(time.Now().Unix()) + "." + sufix
	toFilePath := this.upload + sufix + "/" + fileName
	var size, reSize int64
	path := utils.GetCurrentDir(toFilePath)
	beego.Info("--------path:",path," | toFilePath:",toFilePath)
	if err = this.SaveToFile("file", path); err == nil {
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
	// src, err := Tinify.FromFile(path)
	var reSize int64
	// if err == nil {
	// 	if err = src.ToFile(path); err == nil {
	// 		res, _ := utils.FileSize(path)
	// 		reSize = res
	// 	}
	// }
	// if err != nil {
	// 	beego.Error("compress:", err)
	// }
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

func (this *BaseController) lists(arr ...string) []map[int]interface{} {
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

func (this *BaseController) list(nl [][]NavListItem, arr ...string) {
	var nli [] NavListItem
	for k, v := range arr {
		nli[k].Name = v
		t := k + 1
		if t%10 == 0 {
			nl = append(nl, nli)
			nl = nil
		}
		if k+1 == len(arr) {
			nl = append(nl, nli)
		}
	}
}

// func (this *BaseController) university(index int) []string {
// 	university := []string{"学府", "诗词", "史书典籍", "诗词周边", "美文阅读", "书法欣赏", "其它"}
// 	return university[index:]
// }

// func (this *BaseController) nation(index, end int) []string {
// 	nation := []string{"民族", "阿昌族", "白族", "保安族", "布朗族", "布依族", "藏族", "朝鲜族", "达翰尔族", "傣族", "昂德族", "东乡族", "侗族", "独龙族", "俄罗斯族", "鄂伦春族", "鄂温克族", "高山族", "哈尼族", "哈萨克族", "汉族", "赫哲族", "回族", "基诺族", "京族",
// 		"景颇族", "柯尔克孜族", "拉祜族", "黎族", "傈僳族", "珞巴族", "满族", "毛南族", "门巴族", "蒙古族", "苗族", "仫佬族", "纳西族", "怒族", "普米族", "羌族", "撒拉族", "畲族", "水族", "塔吉克族", "塔塔尔族", "土家族", "图族", "佤族", "维吾尔族", "乌孜别克族", "锡伯族", "瑶族", "彝族", "仡佬族", "裕固族", "壮族"}
// 	if end > 0 {
// 		return nation[index:end]
// 	}
// 	return nation[index:]
// }

// func (this *BaseController) audio(index int) []string {
// 	audio := []string{"音乐", "古琴", "琵琶", "古筝", "笛子", "葫芦丝", "芦笋", "现代流行曲", "古典", "经典"}
// 	return audio[index:]
// }

// func (this *BaseController) video(index int) []string {
// 	video := []string{"视频", "短视频", "长视频", "MV"}
// 	return video[index:]
// }

// func (this *BaseController) original(index int) []string {
// 	original := []string{"原创", "手工艺", "摄影", "纯艺术", "服装", "视频", "音乐"}
// 	return original[index:]
// }

// func (this *BaseController) question(index int) []string {
// 	question := []string{"题库", "英语", "数学", "物理", "化学", "政治", "生物", "地理", "语文", "历史"}
// 	return question[index:]
// }

type Group struct{
	Id int `json:"id"`
	BlogNum string `json:"blogNum"`
	Name string `json:"name"`
	Des string `json:"des"`
	Cover string `json:"cover"`
	BlogCover string `json:"blogCover"`
	Icon string `json:"icon"`
	Official string `json:"official"`
	Format int `json:"format"`
}

func (this *BaseController) group() {
	var groups[] Group 

	img := new(Group)
	img.Format = 0
	img.Name = "图片"
	img.Des = "只显示图片格式的动态"
	img.Cover = "/static/img/group1.png"
	groups = append(groups,*img)

	audio := new(Group)
	audio.Format = 1
	audio.Name = "音频"
	audio.Des="只显示音频格式的动态"
	audio.Cover = "/static/img/group2.png"
	groups = append(groups,*audio)

	video := new(Group)
	video.Format = 2
	video.Name = "视频"
	video.Des = "只显示视频格式的动态"
	video.Cover = "/static/img/group3.png"
	groups = append(groups,*video)

	doc := new(Group)
	doc.Format = 3
	doc.Name = "文字"
	doc.Des="只显示文字格式的动态"
	doc.Cover = "/static/img/group4.png"
	groups = append(groups,*doc)
	this.Data["groups"]=groups
}

func (this *BaseController) banner() {
	uid := this.userId
	maps,id,err := models.Channels(uid,models.CHANNEL_OFFICIAL,[...]interface{}{uid,uid,uid,uid,20,0})
	if err != nil || id==0{			
		this.showTips("data is null")		
	}else{
		this.Data["banner"] = maps
	}
}

func (this *BaseController) qrImg() {
	app := new(models.App)
	res,_ := app.List(1,0)
	beego.Info("app:",app)
	this.Data["android"] = res[0].Qr
}

type Nav struct {
	Name    string
	Action  string
	NavList [] NavList
}

type NavList struct {
	Name        string
	Action      string
	NavListItem [][] NavListItem
}

type NavListItem struct {
	Name   string
	Action string
}

func (this *BaseController) nav() {
	maps,_,_ := models.SqlList(models.CHANNEL_TYPE,[...]interface{}{})
	beego.Info("--------->>>>page:",this.page," | pageSize:",this.pageSize)
	this.Data["nav"] = maps
	this.Data["page"]=this.page
	this.Data["pageSize"] = this.pageSize
}

// func TimeStamp(str string)int64{
// 	beego.Info("-------str:",str)
// 	date,_ := time.Parse("2006-01-02 15:04:05",str)
// 	return date.Unix()
// }

// /**
// * @des 时间转换函数
// * @param atime string 要转换的时间戳（秒）
// * @return string
// */
// func StrTime (atime int64) string{
//     var byTime = []int64{365*24*60*60,24*60*60,60*60,60,1}
//     var unit = []string{"年前","天前","小时前","分钟前","秒钟前"}
//     now := time.Now().Unix()
//     ct := now - atime
//     if ct < 0{
//         return "刚刚"
//     }
//     var res string
//     for i := 0;i < len(byTime);i++{
//         if ct < byTime[i]{
//             continue
//         }
//         var temp = math.Floor(float64(ct / byTime[i]))
//         ct = ct % byTime[i];
//         if temp > 0 {
//             var tempStr string
//             tempStr = strconv.FormatFloat(temp,'f',-1,64)
//             res = MergeString(tempStr,unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
//         }
//         break//我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
//     }
//     return res
// }

// /**
// * @des 拼接字符串
// * @param args ...string 要被拼接的字符串序列
// * @return string
// */
// func MergeString (args ...string) string {
//     buffer := bytes.Buffer{}
//     for i:=0; i<len(args); i++ {
//         buffer.WriteString(args[i])
//     }
//     return buffer.String()
// }

// func ShowNum(str string)string{
// 	num,_ := strconv.Atoi(str);
// 	if num >= 100000000{
// 		return strconv.Itoa(num/100000000.0)+"亿"
// 	}
// 	if num >= 10000 {
// 		return strconv.Itoa(num/10000.0) +"万"		
// 	}
// 	if num >= 1000 {
// 		return strconv.Itoa(num/1000.0)+"千"
// 	}
// 	return strconv.Itoa(num)
// }
