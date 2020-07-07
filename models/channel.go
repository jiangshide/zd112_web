package models

import "github.com/astaxie/beego/orm"
import "github.com/astaxie/beego"

type Channel struct {
	Id int64	`json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	ChannelTypeId int  `json:"typeId"`//频道类型ID
	ChannelNatureId int `json:"natureId"`//频道所属ID
	Cover string `json:"cover"`//封面
	BlogCover string `json:"blogCover"`//封面:最后一个动态的封面
	Name string `json:"name"`//频道名称
	Des string `json:"des"`//频道描述
	Status int `json:"status"` //状态:0~未审核,1~审核中,2~审核通过,-1~移到回忆箱,-2~审核拒绝,-3～禁言，-4~关闭/折叠,-5~被投诉
	Official		int `json:"official"`//官方推荐:-1~取消推荐,0~未推荐,1~推荐,2~特别推荐
	Reason string `json:"reason"`//原由
	Latitude float64 `json:"latitude"`//精度
	Longitude float64 `json:"longitude"`//纬度
	LocationType string `json:"locationType"`//定位类型
}

var CHANNEL_REPLACE = "REPLACE INTO zd_channel(id,uid,channel_nature_id,name,cover,des)VALUES(?,?,?,?,?,?)"

var CHANNEL_ID=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND C.id=? "+CHANNEL_ORDERBY_TIIME//通过频道ID查询:uid,uid,uid,uid,id,limit,offset

var CHANNEL_UID=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND C.uid=? "+CHANNEL_ORDERBY_TIIME//通过用户ID查询:uid,uid,uid,uid,uid,limit,offset
var CHANNEL_UID_ALL=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.uid=? "+CHANNEL_ORDERBY_TIIME//针对用户所有的频道:uid,uid,uid,uid,uid,limit,offset

var CHANNEL_SEARCH = ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND ( C.uid=? OR C.id=? OR instr(C.name,?) >0 ) "+CHANNEL_ORDERBY_TIIME//搜索:uid,uid,uid,uid,uid,id,name,limit,offset

var CHANNEL_FOLLOW=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+" LEFT JOIN zd_user_follow UF ON UF.from_id=C.uid WHERE C.status=2 AND coalesce(UF.status,0) IN(1,2) AND UF.uid=? "+CHANNEL_ORDERBY_TIIME//已关注的:uid,uid,uid,uid,uid,limit,offset
var CHANNEL_UNFOLLOW=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+" LEFT JOIN zd_user_follow UF ON UF.from_id=C.uid WHERE C.status=2 AND coalesce(UF.status,0) NOT IN(1,2) "+CHANNEL_ORDERBY_TIIME//可关注的:uid,uid,uid,uid,limit,offset

var CHANNEL_STATUS=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=? "+CHANNEL_ORDERBY_TIIME//通过状态查询:uid,uid,uid,uid,status,limit,offset

var CHANNEL_LEFT="FROM zd_channel C LEFT JOIN zd_user_profile P ON P.id=C.uid "
var CHANNEL_COMMON=",(SELECT COUNT(1) FROM zd_blog WHERE channel_id=C.id) blogNum,(SELECT name FROM zd_user_remarks WHERE uid=C.uid AND from_id=? ) remark,(SELECT status FROM zd_channel_top WHERE channel_id=C.id AND uid=? ) tops,(SELECT status FROM zd_user_follow WHERE from_id=C.uid AND uid=? ) follows,(SELECT reason FROM zd_report WHERE content_id=C.id AND uid=? ) reportr "
var ChannelSql="SELECT C.id,C.uid,C.channel_type_id typeId,C.channel_nature_id natureId,C.cover,C.blog_cover blogCover,C.name,C.des,C.status,C.official,C.reason,C.update_time date,P.icon,P.nick,P.latitude,P.longitude,P.location_type locationType	 "

var CHANNEL_ORDERBY_TIIME="order by C.update_time DESC limit ? offset ?"

func Channels(uid int64,sql string,ids interface{})(maps *[]orm.Params,id int64,err error){
	maps,id,err = SqlList(sql,ids)
	beego.Info("jsd~maps:",maps,"err:",err," | id:",id," | sql:",sql," | ids:",ids)
	if err != nil || id == 0{
		return
	}
	for _,v := range (*maps){
		res,_,err := Sql(BLOG_CHANNEL, [...]interface{}{uid,uid,uid,uid,uid,v["id"]})
		// beego.Info("err:",err," | res:",res," | ids:",ids)
		if err == nil {
			v["blog"] = res
		// 	files,_,_ := SqlList(FILE,[...]interface{}{v["id"],3,20,0})
		// 	v["urls"] = files
		}
	}
	return
}

func (this *Channel) TableName() string {		
	return TableName("channel")
}

func (this *Channel) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Channel) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Channel) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Channel) Query() error {
	return orm.NewOrm().QueryTable(this.TableName()).Filter("uid",this.Uid).Filter("name",this.Name).One(this)
}

type SingleChannel struct{
	Channel interface{}
	Blog interface{}
}

func InitChannelNature(){
	channelNaturePrivate := new(ChannelNature)
	channelNaturePrivate.Name = "私有频道"
	channelNaturePrivate.Des = "创建属于你的小天地,记录、分享自己的故事"
	_,err :=channelNaturePrivate.Add()
	if err != nil{
		beego.Info(err)
	}
	channelNatureComm := new(ChannelNature)
	channelNatureComm.Name = "公共频道"
	channelNatureComm.Des = "创建有趣或者奇葩的频道,带领大家一起玩"
	_,err = channelNatureComm.Add()
	if err != nil{
		beego.Info(err)	
	}
}

type ChannelNature struct{
	Id int `json:"id"`
	Name string `json:"name"`//频道所属名称
	Des string `json:"des"`//频道所属描述
}

func InitUserChannelNature(uid int64,netInfo,device,ip string,action int){
	channelNature := new(ChannelNature)
	list,err := channelNature.List(20,0)
	beego.Info("---jsd~list:",list," | err:",err)
	beego.Info(list)	
	for _,ct := range list{
		beego.Info(ct.Id)
		channelNatureUser := new(ChannelNatureUser)
		channelNatureUser.ChannelNatureId = ct.Id
		channelNatureUser.Uid = uid
		if(ct.Id == 1){
			channelNatureUser.Num = 1
			channelNatureUser.Time = 1*24*3600
			channelNatureUser.Note = "每天只能创建一个频道"
		}else if(ct.Id == 2){
			channelNatureUser.Num = 1
			channelNatureUser.Time = 2*24*3600
			channelNatureUser.Note = "每两天只能创建一个频道"
		}
		id,err := channelNatureUser.Add()
		beego.Info("---jsd~id:",id," | err:",err)
		if err != nil{
			AddMake(netInfo,device,ip,uid,id,action,err)
		}
	}
}

func (this *ChannelNature) TableName() string {
	return TableName("channel_nature")
}

func (this *ChannelNature) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func (this *ChannelNature) List(pageSize, offSet int) (list []*ChannelNature, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type ChannelNatureUser struct{
	Id int64 `json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	ChannelNatureId int `json:"channelNatureId"`//频道所属ID
	Num int `json:status`//频道所属可创建条数
	Time int `json:"time"`//频道所属可创建的时间周期
	Note string `json:"note"` //标注
}

func (this *ChannelNatureUser) TableName() string {
	return TableName("channel_nature_user")
}

func (this *ChannelNatureUser) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func SqlNature()string{
	return "select CN.id,CN.name,CN.des,NU.channel_nature_id,NU.num,NU.time,NU.note FROM zd_channel_nature CN LEFT JOIN zd_channel_nature_user NU ON NU.channel_nature_id=CN.id where NU.uid=?"
}


type ChannelType struct{
	Id int `json:"id"`
	Name string `json:"name"`//频道类型名称
	Des string `json:"des"`//频道类型描述
	Status int `json:"status"`//状态:1~显示,-1~隐藏
	Sequence int `json:"sequence"`//显示顺序
}

func InitChannelType(){
	types := [...]string{"推荐","热门","直播","明星","萌宝","运动","旅游","家居","美食","艺术","教育","科技","娱乐","电影"}
	for k,v := range types{
		channelType := new(ChannelType)
		channelType.Name = v
		channelType.Status = 1
		channelType.Sequence = k
		_,err := channelType.Add()
		if(err != nil){
			beego.Info(err)
		}
	}
}

var CHANNEL_TYPE_ID=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND C.channel_type_id=? "+CHANNEL_ORDERBY_TIIME//频道所属类型:uid,uid,uid,uid,channel_type_id,limit,offset


func (this *ChannelType) TableName() string {
	return TableName("channel_type")
}

func (this *ChannelType) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func (this *ChannelType) List(pageSize, offSet int) (list []*ChannelType, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

type ChannelTypeUser struct{
	Id int64 `json:"id"`
	Uid   int64  `json:"uid"`//用户ID
	ChannelId int64 `json:"channelId"`//频道ID
	ChannelTypeId int `json:"channelTypeId"`//频道类型ID
	Selected int `json:"selected"`//选中状态:1~选中,0~未选中
	Num int `json:status`//排序顺序
	Note string `json:"note"` //标注
}

func (this *ChannelTypeUser) TableName() string {
	return TableName("channel_type_user")
}

func (this *ChannelTypeUser) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

var CHANNEL_TYPE = "select id,name,des,status,sequence from zd_channel_type order by sequence"
var CHANNEL_TYPE_UID = "select channel_type_id id,name,des,selected,num,note from zd_channel_type_user u,zd_channel_type t where t.id = u.channel_type_id and u.uid=?"

type ChannelTop struct{
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ChannelId   int64  `json:"channelId"`//频道ID
	Status int `json:"status"`//状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶
}

func (this *ChannelTop) TableName()string{
	return TableName("channel_top")
}

func (this *ChannelTop) ReadOrCreates(channelId,uid int64)(created bool, id int64, err error){
	top:=ChannelTop{ChannelId:channelId,Uid:uid}
	created,id,err = orm.NewOrm().ReadOrCreate(&top,"uid","channel_id")
	return
}

type ChannelRecommend struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	ChannelId   int64  `json:"channelId"`//频道ID
	Status int `json:"status"`//状态:-1～取消推荐,0~未推荐,1~推荐,2~特别推荐
	FromId int64 `json:"fromId"`//来自用户ID
}

var CHANNEL_OFFICIAL=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND C.official IN(1,2) "+CHANNEL_ORDERBY_TIIME//官方推荐:uid,uid,uid,uid,limit,offset

var CHANNEL_HOT=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 ORDER BY follows DESC limit ? offset ?"//最多关注的频道频道:uid,uid,uid,uid,limit,offset
var CHANNEL_NEW=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 ORDER BY C.update_time limit ? offset ?"//最新发布的频道:uid,uid,uid,uid,limit,offset
var CHANNEL_MOST=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 ORDER BY C.blogNum limit ? offset ?"//最多动态的频道:uid,uid,uid,uid,limit,offset
var CHANNEL_LIKE=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 ORDER BY C.update_time limit ? offset ?"//最新频道:uid,uid,uid,uid,limit,offset


var CHANNEL_RECOMMEND_UID=ChannelSql+CHANNEL_COMMON+"FROM zd_channel_recommend CR LEFT JOIN zd_channel C ON C.id=CR.channel_id LEFT JOIN zd_user_profile P ON P.id=CR.uid WHERE C.status=2 AND CR.uid=? AND CR.status IN(1,2) "+CHANNEL_ORDERBY_TIIME//针对用户推荐～高~重点处理:uid,uid,uid,uid,uid,limit,offset
var CHANNEL_RECOMMEND_UNFOLLOW=ChannelSql+CHANNEL_COMMON+"FROM zd_channel_recommend CR LEFT JOIN zd_channel C ON C.id=CR.channel_id LEFT JOIN zd_user_profile P ON P.id=CR.uid WHERE C.status=2 AND CR.status NOT IN(1,2) "+CHANNEL_ORDERBY_TIIME//针对已推荐的推荐~中:uid,uid,uid,uid,limit,offset
var CHANNEL_RECOMMEND_ALL=ChannelSql+CHANNEL_COMMON+CHANNEL_LEFT+"WHERE C.status=2 AND C.uid != ? "+CHANNEL_ORDERBY_TIIME//推荐所有~低:uid,uid,uid,uid,uid,limit,offset

func (this *ChannelRecommend) TableName() string {
	return TableName("channel_recommend")
}

func (this *ChannelRecommend) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func SqlChannelFollow(where string)string{
	return "select C.id,C.uid,C.channel_type_id typeId,C.channel_nature_id natureId,C.name,C.des,C.format,C.cover,C.update_time date,UP.nick,CN.name nature,CT.name type from zd_channel C left join zd_user_profile UP on C.uid=UP.id left join zd_channel_nature CN on C.channel_nature_id = CN.id left join zd_channel_type CT on C.channel_type_id = CT.id "+where+" order by C.create_time DESC limit ? offset ?"
}

