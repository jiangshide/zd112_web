package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
	// "fmt"
	"math/rand"
	// "strings"
	"time"
)

type WeiXinToken struct{
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
	Unionid string `json:"unionid"`
}

type WeiXin struct{
	Openid string `json:"openid"`
	NickName string `json:"nickname"`
	Sex int 	`json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	Unionid string 	`json:"unionid"`
}

type WeiXinInfo struct{
	Id         int64 `json:"id"`
	Uid         int64 `json:"uid"`
	Openid string `json:"openid"`
	Nick string `json:"nick"`
	Sex int 	`json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	Icon string `json:"icon"`
	Privilege string `json:"privilege"`
	Unionid string 	`json:"unionid"`
}

func (this *WeiXinInfo) TableName() string {
	return TableName("user_weixin")
}

func (this *WeiXinInfo) Add()(int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *WeiXinInfo) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *WeiXinInfo) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}


type User struct {
	Id         int64 `json:"id"`
	Name  string `json:"name"`//用户名
	Password   string `json:"password"`//密码
	RoleIds    string `json:"roleIds"`//角色id字符串，如：2,3,4
	Salt       string `json:"salt"`//密码盐
	Ip string `json:"ip"`//登录IP
	Status     int `json:"status"` //状态:1~审核中,2~审核通过,3移到回忆箱,-1~审核拒绝,-2~禁言,-3~关闭/折叠,-4~被举报
 	Reason string  `json:"reason"`//原由
 	Online int `json:"online"`//在线状态:1~在线:-1~离线
 	Source int `json:"source"` //注册方式:0~android,1~ios,2~web,3～小程序
}	

func (this *User) TableName() string {
	return TableName("uc_user")
}

func (this *User) GetByName(loginName string) (*User, error) {
	a := new(User)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (this *User) GetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(this.TableName())
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (this *User) GetById(id int64) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(this.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (this *User) Query(key string,value interface{}) error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(key,value).One(this);
	}
	return orm.NewOrm().Read(this)
}

func (this *User) Add() (int64,error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *User) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
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

// func (this *Profile) Query() error {
// 	if this.Id == 0 {
// 		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
// 	}
// 	return orm.NewOrm().Read(this)
// }

func (this *Profile) Query(key string,value interface{}) error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(key,value).One(this);
	}
	return orm.NewOrm().Read(this)
}

func (this *Profile) List(pageSize, offSet int) (list []*Profile, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return	
}

func Location(latitude,longitude float64,locationType string,id int64){
	SqlRaw("UPDATE zd_user_profile SET latitude=?,longitude=?,location_type=? WHERE id=?",[...]interface{}{latitude,longitude,locationType,id})
}

type Active struct{
	Id int64 `json:"active"`
	Uid int64 `json:"uid"`
	ContentId int64 `json:"contentId"`
	Num int `json:"num"`//活跃值
}

func (this *Active) TableName() string {
	return TableName("user_active")
}

func (this *Active) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

//用户修改信息频率
type Frequency struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Type int `json:"type"`
	Number int `json:"number"`
	Frequency int64 `json:"frequency"`
	Time int64 `json:"time"`
}

func InitFrequency(uid,frequency,time int64,types int,number int){
	SqlRaw("INSERT INTO zd_user_frequency(uid,type,number,frequency,time)VALUES(?,?,?,?,?)",[...]interface{}{uid,types,number,frequency,time})
}

var Frequency_UPDATE="UPDATE zd_user_frequency set uid=?,type=?,number=?,frequency=?,time=? WHERE id=?"

func (this *Frequency) TableName() string {
	return TableName("user_frequency")
}

func (this *Frequency) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter("uid",this.Uid).Filter("type",this.Type).One(this)
	}
	return orm.NewOrm().Read(this)
}

//实人认证
type Certification struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Name string `json:"name"`//真实姓名
	IdCard string `json:"idCard"`
	IdCardPicFront string `json:"idCardPicFront"`
	IdCardPicBehind string `json:"idCardPicBehind"`
	Url string `json:"url"`
	Status int `json:"status"`//状态:0~审核中,1~审核通过,-1~审核拒绝
	Reason string `json:"reason"`
}

var CERTIFICATION_STATUS_UID=SqlCertification+"WHERE status=? AND uid=? order by update_time limit ? offset ?"
var CERTIFICATION_STATUS=SqlCertification+"WHERE status=? order by update_time limit ? offset ?"
var CERTIFICATION_LIST=SqlCertification+"order by update_time limit ? offset ?"
var SqlCertification = "SELECT id,uid,name,id_card idCard,id_card_pic_front idCardPicFront,id_card_pic_behind idCardPicBehind,url,status,reason,update_time date FROM zd_user_certification "

var CERTIFICATION_UPDATE = "REPLACE INTO zd_user_certification(id,uid,name,id_card,id_card_pic_front,id_card_pic_behind,url,status,reason)VALUES(?,?,?,?,?,?,?,?,?)"

func (this *Certification) TableName() string {
	return TableName("user_certification")
}

func (this *Certification) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

//备注
type Remarks struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`//用户ID
	FromId int64 `json:"fromId"`//来自用户ID
	Name string `json:"name"`//备注名称
	Url string `json:"url"`
}

func (this *Remarks) TableName() string {
	return TableName("user_remarks")
}

func (this *Remarks) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *Remarks) ReadOrCreates(uid,fromId int64,name,url string)(created bool, id int64, err error){
	remarks:=Remarks{Uid:uid,FromId:fromId,Name:name}
	created,id,err = orm.NewOrm().ReadOrCreate(&remarks,"uid","from_id","name")
	remarks.Url=url
	id,err = remarks.Update()	
	return
}

//历程
type Course struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	Name string `json:"name"`
	Des string `json:"des"`
	Cover string `json:"cover"`
	Url string `json:"url"`
	Source int `json:"source"`//来源:0~未知,1~注册,2~登录,3～创建频道,4~发布动态
	SourceId int64 `json:"sourceId"`
}
	
func CourseAdd(uid,sourceId int64,source int,name,des,cover,url string){
	SqlRaw("INSERT INTO zd_user_course(uid,name,des,cover,url,source,source_id) VALUES (?,?,?,?,?,?,?)",[...]interface{}{uid,name,des,cover,url,source,sourceId})
}

func (this *Course) TableName() string {
	return TableName("user_course")
}

// func (this *Course) Add()(int64,error){
// 	return orm.NewOrm().Insert(this)
// }

// func (this *Course) List(pageSize, offSet int) (list []*Banner, total int64) {
// 	query := orm.NewOrm().QueryTable(this.TableName())
// 	total, _ = query.Count()
// 	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
// 	return
// }

type Invite struct{
	Id              int64 `json:"id"`
	Uid          	int64	`json:"uid"`//用户ID
	FromId 			int64	`json:"fromId"`//内容来源ID
	Times 	int `json:"times"`//邀请次数
	Time 	int `json:"time"`//到期时间
	Status 			int `json:"status"` //状态:0～无效,1～有效,-1:过期
	Especially int `json:"especially"`//有效状态:1~永久,0～取消永久
	Reason  string `json:"reason"`//原由
 	Code	string `json:"code"`//邀请码
}

func (this *Invite) TableName() string {
	return TableName("user_Invite")
}

func (this *Invite) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Invite) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Invite) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Invite) ReadOrCreates(uid,fromId int64)(created bool, id int64, err error){
	invite:=Invite{Uid:uid,FromId:fromId}
	created,id,err = orm.NewOrm().ReadOrCreate(&invite,"uid","from_id")
	return
}

//好友
type Friend struct{
	Id              int64 `json:"id"`
	Uid          	int64	`json:"uid"`
	FromId		int64 `json:"fromId"`
	Status 			int `json:"status"` //状态,0:取消,1~好友,-1:拉黑,-2:删除好友
	Reason	string `json:"reason"`//原由
	Url string `json:"url"`//快照
}

func (this *Friend) TableName() string {
	return TableName("user_friend")
}

func (this *Friend) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Friend) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Friend) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Friend) ReadOrCreates(uid,fromId int64,status int,reason,url string)(created bool, id int64, err error){
	friend:=Friend{Uid:uid,FromId:fromId}
	created,id,err = orm.NewOrm().ReadOrCreate(&friend,"uid","from_id")
	friend.Status=status
	friend.Reason = reason
	friend.Url = url
	id,err = friend.Update()
	return
}

var FRIEND_LIST=SqlUserFriend+"WHERE F.from_id=? AND F.status=? "+FRIEND_ORDER_BY//好友列表
var SqlUserFriend = "SELECT F.id,F.uid,F.status,F.reason,F.url,F.create_time date,P.nick,P.icon,(SELECT COUNT(1) FROM zd_channel WHERE uid=F.uid) channelNum,(SELECT COUNT(1) FROM zd_blog WHERE uid=F.uid) blogNum FROM zd_user_friend F left join zd_user_profile P ON p.id=F.uid "
var FRIEND_ORDER_BY="ORDER BY F.update_time DESC LIMIT ? OFFSET ? "

type UserTop struct{
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	FromId   int64  `json:"fromId"`//内容来源ID:UID
	Status int `json:"status"`//状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶
}

func (this *UserTop) TableName()string{
	return TableName("user_top")
}

func (this *UserTop) ReadOrCreates(fromId,uid int64)(created bool, id int64, err error){
	top:=UserTop{FromId:fromId,Uid:uid}
	created,id,err = orm.NewOrm().ReadOrCreate(&top,"uid","from_id")
	return
}

type UserFollow struct{
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`//用户ID
	FromId int64 `json:"fromId"`//内容来源ID:UID
	Status int `json:"status"`//状态:-1~取消关注,0~推荐关注,1~已关注,2~特别关注
}

var USER_FOLLOWED=UserSql+USER_COMMON+"FROM zd_user_follow F LEFT JOIN zd_user_profile P ON P.id=F.uid LEFT JOIN zd_uc_user U ON U.id=F.uid WHERE F.status IN(1,2) AND F.from_id=? "+USER_ORDER_BY//已关注用户:uid,uid,uid,uid,uid,limit,offset
var USER_FOLLOW=UserSql+USER_COMMON+"FROM zd_user_follow F LEFT JOIN zd_user_profile P ON P.id=F.from_id LEFT JOIN zd_uc_user U ON U.id=F.from_id WHERE F.from_id !=? OR (F.from_id=? AND F.status NOT IN(1,2)) "+USER_ORDER_BY//可关注用户:uid,uid,uid,uid,uid,uid,limit,offset
var USER_FANS=UserSql+USER_COMMON+"FROM zd_user_follow F LEFT JOIN zd_user_profile P ON P.id=F.from_id LEFT JOIN zd_uc_user U ON U.id=F.from_id WHERE F.status IN(1,2) AND F.uid=? "+USER_ORDER_BY//被关注用户:粉丝:uid,uid,uid,uid,uid,limit,offset

func (this *UserFollow) TableName() string {
	return TableName("user_follow")
}

func (this *UserFollow) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *UserFollow) ReadOrCreates()(status int, id int64, err error){
	follow:=UserFollow{Uid:this.Uid,FromId:this.FromId}
	_,id,err = orm.NewOrm().ReadOrCreate(&follow,"uid","from_id")
	follow.Status = 1-follow.Status
	id,err = follow.Update()
	return follow.Status,id,err
}

type UserRecommend struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	UserId   int64  `json:"userId"`//内容来源ID
	Status int `json:"status"`//状态:-1~取消推荐,0~未推荐,1~推荐,2～特别推荐
	FromId int64 `json:"fromId"`//来自用户ID
}

var USER_RECOMMEND=UserSql+USER_COMMON+"FROM zd_user_recommend R LEFT JOIN zd_user_profile P ON R.user_id=P.id LEFT JOIN zd_uc_user U ON R.user_id=U.id WHERE R.uid=? AND R.status IN(1,2) "+USER_ORDER_BY//推荐用户:uid,uid,uid,uid,uid,limit,offset

func (this *UserRecommend) TableName() string {
	return TableName("user_recommend")
}

func (this *UserRecommend) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

// var FOLLOW_WHERE_UID = "where f.from_id=?"
// var FOLLOW_WHERE_STATUS="where f.status=?"

// func SqlFollow(where string)string{
// 	return "select p.name,p.icon,p.sex,p.birthday,p.city_code cityCode,p.update_time date,f.status from zd_uc_user_profile p inner join zd_uc_user_follow f on p.id = f.uid "+where+" order by p.create_time limit ? offset ?"
// }

// var USER ="select id,uid,icon,level,score,vip,name,phone,motto,birthday,sex,constellation,email,weixin,qq,weibo,id_card,id_card_pic_front,id_card_pic_behind,city_code,longitude,latitude,create_time date,(select status from zd_uc_user_friend where uid=? and from_id=?)friend,(select status from zd_uc_user_follow where uid=? and from_id=?)follow from zd_uc_user_profile where uid=?"

// func SqlById(id int64)(user *orm.Params,err error,num int64){
// 	var maps []orm.Params
//  	o := orm.NewOrm()
//    num, err = o.Raw("SELECT * FROM zd_uc_user_profile WHERE uid=?", id).Values(&maps)
//     return &maps[0],err,num
// }

var r *rand.Rand

func init() {
    r = rand.New(rand.NewSource(time.Now().Unix()))
}

func (this *User) GenValidateCode(len int) string {
	// numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	// r := len(numeric)
	// rand.Seed(time.Now().UnixNano())
 
	// var sb strings.Builder
	// for i := 0; i < width; i++ {
	// 	fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	// }
	// return sb.String()

	bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}
