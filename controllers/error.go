package controllers

import "github.com/astaxie/beego"

const (

	MSG_OK           = 0
	MSG_ERR          = -1
	AUTH             = "auth"
	DEFAULT_TIPS     = "该项不能为空!"
	DEFAULT_MIN_SIZE = 1

	//the base
	RES_OK              = 0
	FALSE               = -1
	TOKEN_PRODUCE_FALSE = -2
	TOKEN_INVALIDATE    = -3
	USER_IS_NULL 	= -4
	VALIDATE_FAIL = -5

	//the systen:10~300
	MAC_ISNULL     = -10
	DEVICE_ISNULL  = -11
	GETINT_FALSE   = -12
	GETINT64_FALSE = -13

	//the user:301~500
	DEFAULT_ISNULL    = -301
	DEFAULT_SIZE_LOW  = -302
	PASSWORD_DIFFER   = -303
	USER_NOT_ACTIVE   = -2
	USER_FORBIDDEN    = -1
	USER_FORBIDDEN_WORDS=1
	USER_EXCEPTION=2
	USER_PASSWORD_ERR = -306
	USER_ALREADY_EXIISTA=-307
	USER_SET_INFO=-308

	//the db:501~800
	DB_INSERT_FALSE = 501
	DB_DELETE_FALSE = 502
	DB_UPDATE_FALSE = 503
	DB_QUERY_FALSE  = 503
	//the net:801~1200
	//the other:1201~...

	REG=1
	BIND=11
	LOGIN=2
	CHANNEL=3
	BLOG=4

	DEFAULT_INTRO="平不平凡，从梵记开始..."
)

var msg = map[int]interface{}{
	//the system:10~300
	MAC_ISNULL:    "Mac地址不能为空!",
	DEVICE_ISNULL: "设备名称不能为空!",
	//the user:301~500
	PASSWORD_DIFFER:   "密码不一致!",
	USER_NOT_ACTIVE:   "该账号未激活!",
	USER_FORBIDDEN:    "该账号已被禁用!",
	USER_FORBIDDEN_WORDS:    "该账号已被禁言!",
	USER_EXCEPTION:    "该账号异常，请联系客户!",
	USER_PASSWORD_ERR: "输入密码错误!",
	USER_IS_NULL:"用户不存在!",
	VALIDATE_FAIL:"验证码错误",
	USER_ALREADY_EXIISTA:"用户已存在!",
	USER_SET_INFO:"用户还未设置信息",
	//the db:501~800

	//the net:801~1200

	//the other:1201~...
}

type ErrorController struct{
	beego.Controller
}

func (this *ErrorController)Error404(){
	this.Data["content"]="正在开发中..."
	this.TplName="error/404.html"
}

func (this *ErrorController)Error501(){
	this.Data["content"]="server error"
	this.TplName = "error/501.html"
}

func (this *ErrorController)ErrorDb(){
	this.Data["content"]="database is now down"
	this.TplName = "error/dberror.html"
}