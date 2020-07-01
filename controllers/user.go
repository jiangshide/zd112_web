package controllers

import (
	"zd112_web/models"
	"github.com/astaxie/beego"
	"github.com/jiangshide/GoComm/utils"
	"time"
	// "runtime"
)

// const (
// 	LOGIN = 1
// 	REG   = 2
// )

type UserController struct {
	BaseController
}

func (this *UserController) Login() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		user := new(models.User)
		user.Name = this.getString("username", "账号不能为空!", DEFAULT_MIN_SIZE)
		password := this.getString("password", "密码不能为空!", DEFAULT_MIN_SIZE)
		beego.Info("----------name:",user.Name," | password:",password)
		if err := user.Query("name",user.Name); err != nil {
			beego.Info("err:", err, " | user:", user)
			this.showTips("该账号不存在!")
		}
		if user.Status == USER_FORBIDDEN {
			this.showTips(USER_FORBIDDEN)
		}else if user.Status == -1 ||user.Status == -2 || user.Status == -4{
			this.showTips(USER_EXCEPTION)
		} else if user.Password != utils.Md5ToStr(utils.Md5ToStr(password)+user.Salt) {
			this.showTips(USER_PASSWORD_ERR)
		}else{
			
			this.setCook(user, 10000)
			this.redirect("/")
		}

		// if _, err := user.Update(); err == nil {
		// 	userLocation := new(models.UserLocation)
		// 	userLocation.UserId = user.Id
		// 	this.userId = user.Id
		// 	this.userName = user.Name
		// 	userLocation.Ip = this.getIp()
		// 	userLocation.Mac = this.getMac()
		// 	userLocation.Device = runtime.GOOS
		// 	userLocation.Arch = runtime.GOARCH
		// 	userLocation.SdkVersion = runtime.Version()
		// 	userLocation.AppVersion = this.version
		// 	userLocation.CreateId = this.userId
		// 	userLocation.CreateTime = time.Now().Unix()
		// 	if index, err := userLocation.Add(); err != nil {
		// 		beego.Info("index:", index, " | err:", err)
		// 	}
		// 	this.setCook(user, 10000)
		// 	this.redirect("/")
		// } else {
		// 	this.showTips(err)
		// }
	}
	this.display()
}

func (this *UserController) Reg() {
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		user := new(models.User)
		user.Source = 2
		user.Name = this.getString("username", DEFAULT_TIPS, DEFAULT_MIN_SIZE)
		password := this.getString("password", DEFAULT_TIPS, DEFAULT_MIN_SIZE)
		repassword := this.getString("repassword",DEFAULT_TIPS, DEFAULT_MIN_SIZE)
		if password != repassword{
			this.showTips(PASSWORD_DIFFER)
		}
		// Country := this.getString("country","",0)
		// icon := this.getString("icon",DEFAULT_TIPS,DEFAULT_MIN_SIZE) 
		intro := this.getString("intro","",0)

		// latitude := this.getFloat("latitude",0.0)
		// longitude := this.getFloat("longitude",0.0)
		// locationType := this.getString("locationType",0,0)
		// adCode := this.getString("adCode",0,0)


		// netInfo := this.getString("netInfo",0,0)
		// device:= this.getString("device",0,0)
		
		user.Salt = utils.GetRandomString(10)
		user.Password = utils.Md5ToStr(utils.Md5ToStr(password) + user.Salt)
		user.Status=2
		user.Ip = this.getIp()
		id, err := user.Add();
		if err != nil {
			this.showTips(err)
		}
		profile:=new(models.Profile)
		profile.Id = id
		profile.UnionId = user.Name
		profile.Nick = utils.GetRandomName()
		if intro == ""{
			profile.Intro=DEFAULT_INTRO
		}else{
			profile.Intro=intro
		}
		
		profile.Icon = this.defaultIcon
		profile.Sex = this.getInt("sex",0)
		// profile.Latitude = latitude
		// profile.Longitude = longitude
		// profile.LocationType = locationType
		// profile.Country = country
		// profile.City = position.City
		// profile.AdCode  = adCode
		_,err = profile.Add()
		if err != nil{
			this.showTips(err)		
		}

		models.InitFrequency(id,0,time.Now().Unix(),2,1)
		models.InitFrequency(id,0,time.Now().Unix(),6,1)
		models.InitUserChannelNature(id,"","",this.getIp(),REG)
		models.CourseAdd(id,id,1,"世界，我来了!",DEFAULT_INTRO,this.defaultIcon,"")
		this.redirect("/")
	}
	this.display()
}

func (this *UserController) Forget() {
	beego.ReadFromRequest(&this.Controller)
	this.display()
}

func (this *UserController) Logout() {
	this.Ctx.SetCookie(AUTH, "")
	this.redirect("/")
}
