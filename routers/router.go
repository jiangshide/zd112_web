package routers

import (
	"sanskrit_web/controllers"
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

	beego.Router("/add", &controllers.AppController{}, "*:Add")

	beego.Router("/blog",&controllers.BlogController{})
	beego.Router("/blog/ajaxFollow", &controllers.BlogController{}, "*:AjaxFollow")
	beego.Router("/blog/ajaxPraise", &controllers.BlogController{}, "*:AjaxPraise")
	beego.Router("/blog/ajaxCommendAdd", &controllers.BlogController{}, "*:AjaxCommendAdd")
	
		
}
