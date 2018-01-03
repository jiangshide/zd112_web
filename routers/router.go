package routers

import (
	"zd112_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/user/reg", &controllers.UserController{}, "*:Reg")

	beego.Router("/university",&controllers.UniversityController{})

	beego.Router("/upload", &controllers.BaseController{}, "*:Upload")
	beego.ErrorController(&controllers.ErrorController{})
}
