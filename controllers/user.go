package controllers

import (
	"zd112_web/models"
	"github.com/astaxie/beego"
	"zd112/utils"
	// "time"
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
		if err := user.Query(); err != nil {
			beego.Info("err:", err, " | user:", user)
			this.showTips("该账号不存在!")
		}
		// if user.Status == 0 {
		// 	this.showTips("该账号未激活!")
		// } else if user.Status == 2 {
		// 	this.showTips("该账号已被禁用!")
		// } else if user.Password != utils.Md5(password+user.Salt) {
		// 	this.showTips("该账号密码错误!")
		// }
		if user.Status == USER_FORBIDDEN {
			this.showTips(USER_FORBIDDEN)
		}else if user.Status == -1 ||user.Status == -2 || user.Status == -4{
			this.showTips(USER_EXCEPTION)
		} else if user.Password != utils.Md5(utils.Md5(password)+user.Salt) {
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
		user.Name = this.getString("username", "账号不能为空!", DEFAULT_MIN_SIZE)
		password := this.getString("password", "密码不能为空!", DEFAULT_MIN_SIZE)
		rePassword := this.getString("repassword", "请再次输入密码!", DEFAULT_MIN_SIZE)
		if password != rePassword {
			this.showTips("密码不一致!")
		}
		user.Salt = utils.GetRandomString(10)
		user.Password = utils.Md5(password + user.Salt)
		if _, err := user.Add(); err != nil {
			this.showTips(err)
		}
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
