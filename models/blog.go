package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"
 "encoding/json"
 "github.com/jiangshide/GoComm/utils"
 )

type Blog struct {
	Id int64	`json:"id"`
	Uid  int64  `json:"uid"`//用户ID
	At string `json:"at"`//@用户
	ChannelId int64 `json:"channelId"`//频道ID
	PositionId int64 `json:"positionId"`//当前位置ID
	Content string `json:"content"`//内容
	Title string `json:"title"`//动态名称
	Des string `json:"des"`//动态描述
	Latitude float64 `json:"latitude"`//精度
	Longitude float64 `json:"longitude"`//纬度
	LocationType string `json:"locationType"`//定位类型
	Country string `json:"country"`//国家
	City string `json:"city"`//城市
	Position string `json:"position"` //位置
	Address string `json:"address"`//详细位置
	CityCode string `json:"cityCode"`//城市编码
	AdCode string `json:"adCode"`	//区域码 
	TimeZone string `json:"timeCone"`//时区
	Tag string `json:"tag"`//标签
	Status int `json:"status"`//状态:0~未审核,1~审核中,2~审核通过,-1~移到回忆箱,-2~审核拒绝,-3～禁言，-4~关闭/折叠,-5~被投诉
	Reason string `json:"reason"` //原由
	Official		int `json:"status"`//官方推荐:-1~取消推荐,0~未推荐,1~推荐,2~特别推荐
	Cover string `json:"cover"`//封面
	Url string `json:"url"`//文件Url
	Name string `json:"name"`//文件名称
	Sufix string `json:"sufix"`//文件名后缀
	Format int `json:"format"`//内容格式:0:图片,1:音频,2:视频,3:文档,4:web,5:VR
	Duration int64 `json:"duration"` //内容时长
	Width int `json:"width"`//内容宽
	Height int `json:"height"`//内容高
	Size int64 `json:"size"`//内容尺寸
	Rotate int `json:"rotate"`//角度旋转
	Bitrate int `json:"bitrate"`//采用率
	SampleRate int `json:"SampleRate"`//频率
	Level int `json:"Level"`//质量:0~标准
	Mode int `json:"mode"`//模式
	Wave string `json:"wave"`//频谱
	LrcZh string `json:"lrcZh"`//字幕~中文
	LrcEs string `json:"lrcEn"`//字母~英文
	Source int `json:"source"`//创作类型:0~原创,1~其它,5~本地同步			
}

var BLOG_CHANNEL_NEW=BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_channel C ON C.id=B.channel_id WHERE B.status=2 AND C.id=? "+BLOG_SORT_DATE//针对频道最新:uid,uid,uid,uid,uid,channel_id,limit,offset
var BLOG_CHANNEL_HOT=BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_channel C ON C.id=B.channel_id WHERE B.status=2 AND C.id=? "+ORDER_BY_PRAISE//针对频道最热:uid,uid,uid,uid,uid,channel_id,limit,offset

var BLOG_HISTORY = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_history H ON H.content_id=B.id WHERE H.uid=? AND H.source=? AND B.status=2 ORDER BY H.update_time DESC LIMIT ? OFFSET ?"//查看历史记录:uid,uid,uid,uid,uid,uid,source

var BLOG_FOLLOW = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_user_follow UF ON UF.uid=B.uid AND UF.uid=UP.id WHERE (UF.from_id=? AND coalesce(UF.status,0) IN(1,2)) AND B.status=2 AND ZUF.status >= 0 "+BLOG_SORT_DATE//关注:uid,uid,uid,uid,uid,uid,limit,offset
var BLOG_FIND = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_blog_recommend BR ON BR.blog_id = B.id LEFT JOIN zd_user_follow UF ON UF.uid=B.uid WHERE B.uid!=? AND B.status=2 AND coalesce(BR.status,0) = 0 AND coalesce(UF.status,0) = 0 "+BLOG_SORT_DATE//发现:uid,uid,uid,uid,uid,uid,limit,offset
var BLOG_HOMETOWN = BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE UP.city=? AND B.status=2 "+BLOG_SORT_DATE//通过用户出生城市查询~老乡:uid,uid,uid,uid,uid,city,limit,offset
var BLOG_CITY = BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.city=? AND B.status=2 "+BLOG_SORT_DATE//通过城市查询:uid,uid,uid,uid,uid,city,limit,offset
var BLOG_UID = BlogSql+BLOG_COMMON+BLOG_USER_TOP+BLOG_LEFT+"WHERE B.uid=? AND B.status=2 "+BLOG_SORT_DATE//通过用户ID查询:uid,uid,uid,uid,uid,uid,uid,limit,offset
var BLOG_ID =BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.status=2 AND B.id=? "//通过id查询:uid,uid,uid,uid,uid,id
var BLOG_CHANNEL =BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.status=2 AND B.channel_id=? "//通过channel_id查询:uid,uid,uid,uid,uid,channel_id
var BLOG_NAME_SUFIX =BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.status=2 AND B.uid=? AND B.name=? AND B.sufix=? "//通过名称后缀查询:uid,uid,uid,uid,uid,uid,name,sufix

var BLOG_FORMAT=BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.status=2 AND B.format=? "+BLOG_SORT_DATE//通过格式查询:uid,uid,uid,uid,uid,format
var BLOG_STATUS =BlogSql+BLOG_COMMON+BLOG_LEFT+"WHERE B.status=? "+BLOG_SORT_DATE//通过状态查询:uid,uid,uid,uid,uid,status
var BLOG_LEFT="FROM zd_blog B LEFT JOIN zd_user_profile UP ON B.uid=UP.id LEFT JOIN zd_user_friend ZUF ON ZUF.uid=B.uid "

var BLOG_COMMON = ",(SELECT name FROM zd_channel WHERE id=B.channel_id ) channel,(SELECT COUNT(1) FROM zd_blog_praise WHERE blog_id=B.id AND status IN(1,2) ) praiseNum,(SELECT COUNT(1) FROM zd_blog_view WHERE blog_id=B.id ) viewNum,(SELECT COUNT(1) FROM zd_blog_share WHERE blog_id=B.id AND status=1 ) shareNum,(SELECT COUNT(1) FROM zd_comment WHERE content_id=B.id)+(SELECT COUNT(1) FROM zd_comment_uid WHERE content_id=B.id) commentNum,(SELECT name FROM zd_user_remarks WHERE uid=B.uid AND from_id=? ) remark,(SELECT status FROM zd_blog_praise WHERE blog_id=B.id AND uid=? ) praises,(SELECT reason FROM zd_report WHERE content_id=B.id AND source=3 AND uid=? ) reportr,(SELECT status FROM zd_user_follow WHERE uid=B.uid AND from_id=? ) follows,(SELECT status FROM zd_blog_collection WHERE blog_id=B.id AND uid=? ) collections "

var BLOG_USER_TOP=",(SELECT status FROM zd_blog_top WHERE blog_id=B.id AND uid=? ) tops "

var BlogSql = "SELECT B.id,B.uid,B.at,B.channel_id channelId,B.position_id positionId,B.content,B.title,B.des,B.latitude,B.longitude,B.location_type locationType,B.country,B.city,B.position,B.address,B.city_code cityCode,B.ad_code adCode,B.time_zone timeCone,B.tag,B.status,B.reason,B.official,B.url,B.cover,B.name,B.sufix,B.format,B.duration,B.width,B.height,B.size,B.rotate,B.bitrate,B.sample_rate sampleRate,B.level,B.mode,B.wave,B.lrc_zh lrcZh,B.lrc_es lrcEs,B.source,B.create_time date,UP.icon,UP.nick,UP.sex,UP.age,UP.zodiac,UP.city ucity "

var BLOG_SORT_DATE ="ORDER BY B.create_time DESC LIMIT ? OFFSET ? "//按时间


var ORDER_BY_TIME="ORDER BY B.create_time DESC LIMIT ? OFFSET ? "//按时间
var ORDER_BY_PRAISE="ORDER BY praiseNum DESC LIMIT ? OFFSET ?"//按点赞

//删除	
var BLOG_DELETE = "DELETE zd_blog,zd_blog_file,zd_blog_praise,zd_blog_collection,zd_comment,zd_blog_recommend,zd_blog_top,zd_blog_follow,zd_blog_view,zd_blog_share,zd_blog_report FROM zd_blog LEFT JOIN zd_blog_file ON zd_blog_file.blog_id=zd_blog.id LEFT JOIN zd_blog_praise ON zd_blog_praise.blog_id=zd_blog.id LEFT JOIN zd_comment ON zd_comment.blog_id=zd_blog.id LEFT JOIN zd_blog_recommend ON zd_blog_recommend.blog_id=zd_blog.id LEFT JOIN zd_blog_top ON zd_blog_top.blog_id=zd_blog.id LEFT JOIN zd_blog_follow ON zd_blog_follow.blog_id=zd_blog.id LEFT JOIN zd_blog_view ON zd_blog_view.blog_id=zd_blog.id LEFT JOIN zd_blog_share ON zd_blog_share.blog_id=zd_blog.id LEFT JOIN zd_blog_report ON zd_blog_report.blog_id=zd_blog.id LEFT JOIN zd_blog_collection ON zd_blog.id=zd_blog_collection.blog_id WHERE zd_blog.id=?"
	
func (this *Blog) TableName() string {
	return TableName("blog")
}

func (this *Blog) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}

func (this *Blog) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *Blog) Query() error {
	return orm.NewOrm().QueryTable(this.TableName()).Filter("uid",this.Uid).Filter("name",this.Name).Filter("sufix",this.Sufix).One(this)
}

func (this *Blog) ReadOrCreates()(created bool, id int64, err error){
	blog:=Blog{Uid:this.Uid,Name:this.Name,Sufix:this.Sufix}
	created,id,err = orm.NewOrm().ReadOrCreate(&blog,"uid","name","sufix")
	this.Id=id
	this.LrcZh = GetLrcStr(id,this.Name)
	id,err = this.Update()
	return
}

func Blogs(uid int64,sql string,ids interface{})(maps *[]orm.Params,id int64,err error){
	maps,id,err = SqlList(sql,ids)
	beego.Info("err:",err," | id:",id)
	if err != nil || id == 0{
		return
	}
	for _,v := range (*maps){
		files,_,_ := SqlList(FILE,[...]interface{}{v["id"],3,20,0})
		v["urls"] = files
		comments,_,_ := SqlList(COMMENT_STATUS,[...]interface{}{uid,uid,uid,v["id"],2,0})
		for _,v := range *comments{
			v["date"] = utils.StrTime(utils.TimeStamp(v["date"].(string)))
		}
		v["comments"] = comments
		res,_,_ := Sql("SELECT id,position,size,txt_color,bg_color,scroll FROM zd_blog_style WHERE blog_id=?",[...]interface{}{v["id"]})
		v["style"]=res
	}
	return
}

type BlogStyle struct{
	Id int64 `json:"id"`
	BlogId int64 `json:"blogId"`//动态ID
	Position int `json:"position"`//位置:左上-51,上中-49,上右-53,右中-21,右下-85,下中-81,下左-83,左中-19
	Size int `json:"size"`//文字大小
	TxtColor string `json:"txtColor"`//文字颜色
	BgColor string `json:"bgColor"`//背景颜色
	Scroll int `json:"scroll"`//是否滚动:1~滚动,0~不滚动
}


func (this *BlogStyle) TableName() string {		
	return TableName("blog_style")
}

func (this *BlogStyle) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func AddBlogStyle(styleJson string,id int64){
	beego.Info("styleJson:",styleJson)
	var blogStyle BlogStyle
	err := json.Unmarshal([]byte(styleJson),&blogStyle)
	beego.Info("err:",err)
	blogStyle.BlogId = id
	blogStyle.Add()
}


type BlogPraise struct{
	Id int64	`json:"id"`
	Uid int64	`json:"uid"` //用户ID
	BlogId int64 `json"blog_id"` //动态ID
	Status int `json:"status"`//状态:-1~取消,1~点赞,2~超级赞
}

var BLOG_PRAISE = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_blog_praise BP ON BP.blog_id = B.id AND BP.uid=B.uid WHERE BP.uid=? AND coalesce(BP.status,0) IN(1,2) "+BLOG_SORT_DATE//点赞:uid,uid,uid,uid,uid,uid,limit,offset

func (this *BlogPraise) TableName() string {		
	return TableName("blog_praise")
}

func (this *BlogPraise) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *BlogPraise) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *BlogPraise) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

// func (this *BlogPraise) Query() error {
// 	return orm.NewOrm().QueryTable(this.TableName()).Filter("id",this.Id).Filter("uid",this.Uid).Filter("blog_id",this.BlogId).One(this)
// }

func (this *BlogPraise) ReadOrCreates(uid,blogId int64,status int)(created bool, id int64, err error){
	praise:=BlogPraise{Uid:uid,BlogId:blogId}
	created,id,err = orm.NewOrm().ReadOrCreate(&praise,"uid","blog_id")
	praise.Status=status
	id,err = praise.Update()
	return
}

type BlogRecommend struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	BlogId   int64  `json:"blogId"`//动态ID
	Status int `json:"status"`//状态:-1～取消推荐,0~未被推荐,~1推荐,2~特别推荐
	FromId int64 `json:"fromId"`//来自用户ID
}

func (this *BlogRecommend) TableName() string {		
	return TableName("blog_recommend")
}

var BLOG_RECOMMEND_ADD = "REPLACE INTO zd_blog_recommend(id,uid,blog_id,status)VALUES(?,?,?,?)"//添加推荐
var BLOG_RECOMMEND =BlogSql+BLOG_COMMON+"FROM zd_blog_recommend BR LEFT JOIN zd_blog B ON BR.blog_id = B.id LEFT JOIN zd_user_profile UP ON (BR.uid=UP.id AND BR.blog_id=B.id) WHERE BR.uid=? AND BR.status in(1,2) "+BLOG_SORT_DATE//推荐:uid,uid,uid,uid,uid,uid,limit,offset
var BLOG_RECOMMEND_FORMAT =BlogSql+BLOG_COMMON+"FROM zd_blog_recommend BR LEFT JOIN zd_blog B ON BR.blog_id = B.id LEFT JOIN zd_user_profile UP ON (BR.uid=UP.id AND BR.blog_id=B.id) WHERE BR.uid=? AND B.format=? AND BR.status in(1,2) "+BLOG_SORT_DATE//推荐:uid,uid,uid,uid,uid,uid,format,limit,offset



type BlogTop struct{
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	BlogId   int64  `json:"blogId"`//动态ID
	Status int `json:"status"`//状态:-1~取消置顶,0~未置顶,1~置顶,2~超级置顶
}

func (this *BlogTop) TableName() string {		
	return TableName("blog_top")
}

func (this *BlogTop) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *BlogTop) ReadOrCreates()(created bool, id int64, err error){
	top := BlogTop{Uid:this.Uid,BlogId:this.BlogId}
	created,id,err = orm.NewOrm().ReadOrCreate(&top,"uid","blog_id")
	top.Status=this.Status
	id,err = top.Update()
	return
}

type BlogView struct{
	Id int64 `json:"id"`
	Uid   int64  `json:"uid"`//来自于用户ID
	BlogId int64 `json:"blogId"`//内容来源ID
	Num int `json:""mum`//内容详情位置:来自哪张图片
}

var BLOG_BROWSE = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_blog_view BV ON BV.blog_id = B.id AND BV.uid=B.uid WHERE BV.uid=? "+BLOG_SORT_DATE//浏览:uid,uid,uid,uid,uid,uid,limit,offset

func (this *BlogView) TableName() string {		
	return TableName("blog_view")
}

type BlogShare struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	BlogId   int64  `json:"blogId"`//动态ID
	Status int `json:"status"`//状态:1~成功,0～失败
	FromId int64 	`json:"fromId"`//来自用户ID
	Station int `json:"station"`//类型:0~站内,1～微信朋友圈,2~微信好友...
}

func (this *BlogShare) TableName() string {		
	return TableName("blog_share")
}

var BLOG_SHARE = "INSERT INTO zd_blog_share(uid,blog_id,status,from_id,station)VALUES(?,?,?,?,?)"//分享
var BLOG_UPDATE_SHARE = "UPDATE zd_blog SET share=? WHERE id=?"//更新分享条数

func ShareAdd(uid,blogId,fromId int64,status,station int) (count interface{},err error){
	_,err = SqlRaw(BLOG_SHARE,[...]interface{}{uid,blogId,status,fromId,station})
	if err != nil{
		return
	}
	maps := make(map[string]interface{})
	maps["blog_id"]=blogId
	maps["status"]=1//成功的状态
	count,err = SqlCount("zd_blog_share",maps)
	if err != nil{
		return
	}
	_,err = SqlRaw(BLOG_UPDATE_SHARE,[...]interface{}{count,blogId})
	return 
}	

type BlogCollection struct {
	Id int64	`json:"id"`
	Uid int64 `json:"uid"`//用户ID
	BlogId   int64  `json:"blogId"`//动态ID
	Status int `json:"status"`//状态:-1~取消收藏,0～未收藏,1~收藏,2~珍藏
}

func (this *BlogCollection) TableName() string {		
	return TableName("blog_collection")
}

func (this *BlogCollection) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}	

func (this *BlogCollection) ReadOrCreates()(created bool, id int64, err error){
	collection := BlogCollection{Uid:this.Uid,BlogId:this.BlogId}
	created,id,err = orm.NewOrm().ReadOrCreate(&collection,"uid","blog_id")
	collection.Status=this.Status
	id,err = collection.Update()
	return
}

var BLOG_COLLECTION = BlogSql+BLOG_COMMON+BLOG_LEFT+"LEFT JOIN zd_blog_collection BC ON BC.blog_id = B.id WHERE BC.uid=? AND coalesce(BC.status,0) IN(1,2) "+BLOG_SORT_DATE//收藏:uid,uid,uid,uid,uid,uid,limit,offset


