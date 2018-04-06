package routers

import (
	"zd112_web/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/test", &controllers.TestController{})

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/user/reg", &controllers.UserController{}, "*:Reg")
	beego.Router("/user/forget", &controllers.UserController{}, "*:Forget")
	beego.Router("/user/logout", &controllers.UserController{}, "*:Logout")

	beego.Router("/university", &controllers.UniversityController{})
	beego.Router("/university/detail", &controllers.UniversityDetailController{})
	beego.Router("/nation", &controllers.NationController{})
	beego.Router("/audio", &controllers.AudioController{})
	beego.Router("/video", &controllers.VideoController{})

	beego.Router("/upload", &controllers.BaseController{}, "*:Upload")
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.ErrorController(&controllers.ErrorController{})
}
