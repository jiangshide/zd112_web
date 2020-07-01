package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"	
 "encoding/json"
 "fmt"
 "errors")

type Device struct{
	Id         int64
	Uuid	string //获取唯一设备号	
}

func (this *Device) TableName() string {
	return TableName("uc_user_device")
}

func (this *Device) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

type DeviceInfo struct {
	Id         int64
	DeviceId int64 //设备ID
	Uid        int64//用户ID
	Uuid	string //获取唯一设备号	
	PositionId int64//位置ID
	NetOperator string //获取运营商
	NetName string //获取联网方式
	NetSpeed string////网络速度
	Meid string //获取meid
	Imei1 string //获取imei1
	Imei2 string //获取imei2
	Inumeric string //获取双卡手机的imei
	TotalMem int64//总共内存
	Threshold int64//内存阀值
	AvailMem int64 //可用内存
	Mac        string
	Board string//主板
	Brand		string //品牌
	Device string //设备参数
	Display string //显示屏参数
	Fingerprint string //唯一编号
	Serial string //硬件序列号
	Manufacturer string //硬件制造商
	Model       string //机型
	Hardware string//硬件名
	Product string//手机产品名
	Type string// Builder类型
	Host string// HOST值
	User string // User名
	Time string// 编译时间
	OsVersion string//os版本号
	OsName string//os名称
	OsArch string//os架构
	SdkVersion int//当前sdk版本
	AppName	string //应用名称
	Pkg	string //包名
	VersionCode  int//应用code
	VersionName string//应用版本号
	Os string //操作系统
	Resolution string //分辨率800x400
	TimeZone string //时区
	Battery int//电量
	ElapsedRealtime int64//开机时长
	Language string//语言
}

func (this *DeviceInfo) TableName() string {
	return TableName("uc_user_device_info")
}

func (this *DeviceInfo) Add() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *DeviceInfo) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *DeviceInfo) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

func (this *DeviceInfo) Query() error {
	if this.Id == 0 {
		return orm.NewOrm().QueryTable(this.TableName()).Filter(Field(this)).One(this)
	}
	return orm.NewOrm().Read(this)
}

func (this *DeviceInfo) List(pageSize, offSet int) (list []*DeviceInfo, total int64) {
	query := orm.NewOrm().QueryTable(this.TableName())
	total, _ = query.Count()
	query.OrderBy("-id").Limit(pageSize, offSet).All(&list)
	return
}

func AddMake(netInfoJson,deviceJson,ip string,uid,fromId int64,event int,err error)*Position{
	beego.Info("---------netInfoJson:",netInfoJson)
	position := new(Position)
	json.Unmarshal([]byte(netInfoJson),&position)
	beego.Info("--------position:",position)
	position.Ip = ip
	position.Uid = uid
	position.FromId = fromId
	position.Event = event
	if err != nil{
		position.Err = fmt.Sprintf("%s",err)
	}else{
		position.Err = fmt.Sprintf("%s",errors.New("success"))
	}
	id,positionErr := position.Add()	
	if positionErr != nil{
		beego.Info("positionErr:",positionErr)
	}

	deviceInfo := new(DeviceInfo)
	json.Unmarshal([]byte(deviceJson),&deviceInfo)

	device := new(Device)
	device.Uuid = deviceInfo.Uuid
	deviceId,deviceErr := device.Add()
	if deviceErr != nil{
		beego.Info("deviceErr:",deviceErr)
	}

	deviceInfo.DeviceId = deviceId
	deviceInfo.Uid =uid
	deviceInfo.PositionId = id
	_,deviceInfoErr := deviceInfo.Add()
	if deviceErr != nil{
		beego.Info("deviceInfoErr:",deviceInfoErr)	
	}

	return position
}

type Position struct{
	Id int64 `json:"id"`
	Uid        int64//用户ID
	FromId int64//来自ID
	Event int//事件
	Err string//错误描述
	Ip string `json:"ip"`
	Latitude float64 `json:"latitude"`//精度
	Longitude float64 `json:"longitude"`//纬度
	LocationType string `json:"locationType"`//定位类型
	Accuracy string `json:"accuracy"` //精度
	Provider string `json:"provider"`//提供者
	Speed string `json:"Speed"` //速度
	Bearing string `json:"Bearing"`//角度
	Satellites string `json:"Satellites"`//星数
	Country string `json:"country"`//国家
	Province string `json:"province"`//省
	City string `json:"city"`//市 
	District string `json:"district"`//区 
	CityCode string `json:"cityCode"` //城市编码
	AdCode string `json:"adCode"`	//区域码 
	Address string `json:"address"`//地址	
	PoiName string `json:"poiName"`//兴趣点
	NetworkType string `json:"networkType"`//网络类型
	GPSStatus string `json:"GPSStatus"`//GPS状态
	GPSSatellites string `json:"GPSSatellites"`//GPS星数
	TimeZone	string `json:timeZone`//时区
}

func (this *Position) TableName() string {
	return TableName("uc_position")
}	

func (this *Position) Add() (int64,error)  {
	return orm.NewOrm().Insert(this)
}