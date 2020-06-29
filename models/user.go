package models

import "github.com/astaxie/beego/orm"

type User struct {
	// Id         int64
	// Name       string
	// Password   string
	// Salt       string
	// Status     int
	// CreateId   int64
	// UpdateId   int64
	// CreateTime int64
	// UpdateTime int64
	Id         int64 `json:"id"`
	Name  string `json:"name"`//用户名
	Password   string `json:"password"`//密码
	RoleIds    string `json:"roleIds"`//角色id字符串，如：2,3,4
	Salt       string `json:"salt"`//密码盐
	Ip string `json:"ip"`//登录IP
	Status     int `json:"status"` //状态:1~审核中,2~审核通过,3移到回忆箱,-1~审核拒绝,-2~禁言,-3~关闭/折叠,-4~被举报
 	Reason string  `json:"reason"`//原由
 	Online int `json:"online"`//在线状态:1~在线:-1~离线
}

func (this *User) TableName() string {
	return TableName("uc_user")
}

func (this *User) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *User) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *User) List(pageSize, offSet int) (list []*User, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

//用户基本信息
type Profile struct {
	Id              int64 `json:"id"`
	UnionId string `json:"unionId"`//唯一凭证:针对第三登录
	Icon            string `json:"icon"` //头像	
	Level           int `json:"level"`//级别
	Score           int `json:"score"`//梵豆
	Vip				int `json:"vip"`//vip:0~不是vip,1~vip级别1,2～vip级别2....
	Certification int64 `json:"certification"`//实名认证ID
	Nick			string `json:"nick"`//昵称
	FullName            string `json:"fullName"`//姓名
	Phone           string `json:"phone"`//电话号码
	Intro           string `json:"intro"`//座右铭
	Birthday		string `json:"birthday"`//生日	
	Sex             int 	`json:"sex"`//性别:0:保密,1:男,2:女
	Age int `json:"age"`//年龄
	Nation string `json:"nation"`// 民族
	Zodiac string `json:"zodiac"`//生肖
	Constellation	string `json:"constellation"`//星座
	Hobby string `json:"hobby"`//爱好
	Email           string `json:"email"`//邮箱
	Weixin          string `json:"weixin"`//微信
	Qq              string `json:"qq"`//QQ号
	Weibo           string `json:"weibo"`//微博
	Alipay			string `json:"alipay"`//支付宝
	Latitude float64 `json:"latitude"`//精度
	Longitude float64 `json:"longitude"`//纬度
	LocationType string `json:"locationType"`//定位类型
	Country		string `json:"country"`//国家代码
	Province string `json:"province"`//省
	City 		string `json:"city"`//城市
	AdCode string `json:"adCode"`	//区域码 

	PraiseNotice int `json:"praiseNotice"`//点赞通知:0~打开,-1~关闭
	FollowNotice int `json:"followNotice"`//关注通知:0~打开,-1~关闭
	BroweHomeNotice int `json:"broweHomeNotice"`//看主页通知:0~打开,-1~关闭
	BroweChannelNotice int `json:"broweChannelNotice"`//看频道通知:0~打开,-1~关闭
	BroweBlogNotice int `json:"broweBlogNotice"`//看动态通知:0~打开,-1~关闭
	CreateChannelNotice int `json:"createChannelNotice"`//创建频道通知:0~打开,-1~关闭
	CreateBlogNotice int `json:"createBlogNotice"`//发布动态通知:0~打开,-1~关闭
	NoticeType string `json:"noticeType"`//通知类型:可选范围为 -1～7 ，对应 Notification.DEFAULT_ALL = -1 或者 Notification.DEFAULT_SOUND = 1, Notification.DEFAULT_VIBRATE = 2, Notification.DEFAULT_LIGHTS = 4 的任意 “or” 组合。默认按照 -1 处理'
}

var USER_ONLINE = USER_STATUS+"add P.online=? "+USER_ORDER_BY//在线状态:fromId,fromId,uid,uid,status,online,limit,offset

var USER_STATUS=UserSql+USER_COMMON+USER_LEFT+"WHERE U.status=? "+USER_ORDER_BY//状态用户:fromId,fromId,uid,uid,status,limit,offset

var USER_ID=UserSql+USER_COMMON+USER_LEFT+"WHERE P.id=? "//通过id查询:fromId,fromId,uid,uid,id

var USER_LEFT = "FROM zd_user_profile P LEFT JOIN zd_uc_user U ON U.id=P.id "
var USER_COMMON=",(SELECT COUNT(1) FROM zd_user_follow WHERE uid=P.id )fansNum,(SELECT COUNT(1) FROM zd_user_follow WHERE from_id=P.id ) followsNum,(SELECT COUNT(1) FROM zd_channel WHERE uid=P.id) channelsNum,(SELECT COUNT(1) FROM zd_blog WHERE uid=P.id ) blogsNum,(SELECT SUM(num) FROM zd_user_active WHERE uid=P.id ) activeNum,(SELECT name FROM zd_user_remarks WHERE uid=P.id AND from_id=? ) remark,(SELECT status FROM zd_user_friend WHERE uid=P.id AND from_id=? ) friends,(SELECT status FROM zd_user_follow WHERE from_id=P.id AND uid=?) follows,(SELECT reason FROM zd_report WHERE content_id=P.id AND uid=? ) reportr "

var UserSql = "SELECT P.id,P.union_id unionId,P.icon,P.level,P.score,P.vip,P.certification,P.nick,P.full_name fullName,P.phone,P.intro,P.birthday,P.sex,P.age,P.nation,P.zodiac,P.constellation,P.hobby,P.email,P.weixin,P.qq,P.weibo,P.alipay,P.latitude,P.longitude,P.location_type locationType,P.country,P.province,P.city,P.ad_code adCode,P.praise_notice praiseNotice,P.follow_notice followNotice,P.browe_home_notice broweHomeNotice,P.browe_channel_notice broweChannelNotice,P.browe_blog_notice broweBlogNotice,P.create_channel_notice createChannelNotice,P.create_blog_notice createBlogNotice,P.notice_type noticeType,P.create_time date,U.name,U.status,U.reason,U.online "

var USER_ORDER_BY="ORDER BY P.update_time DESC LIMIT ? OFFSET ? "

func (this *Profile) TableName() string {
	return TableName("user_profile")
}

func (this *Profile) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Profile) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Profile) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Profile) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *Profile) List(pageSize, offSet int) (list []*Profile, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return	
}

