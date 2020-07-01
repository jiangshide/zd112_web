package models

import ("github.com/astaxie/beego/orm"
 "github.com/astaxie/beego"	
 "encoding/json"
 )

type File struct{
	Id int64 `json:"id"`
	ContentId int64 `json:"contentId"`//内容来源ID
	Classify int `json:"classify"`//内容分类:1~用户,2~频道,3~动态,4~评论
	Cover string `json:"cover"`//文件封面
	Url string `json:"url"`//内容Url
	Name string `json:"name"`//内容名称
	Sufix string `json:"sufix"`//文件名后缀
	Format int `json:"format"`//内容格式:0:图片,1:音频,2:视频,3:文档,4:web,5:VR
	Duration int64 `json:"duration"` //内容时长
	Width int `json:"width"`//内容宽
	Height int `json:"height"`//内容高
	Size int64 `json:"size"`//内容尺寸
	Rotate int `json:"rotate"`//角度旋转
	Bitrate int `json:"bitrate"`//采用率
	SampleRate int `json:"sample_rate"`//频率
	Level int `json:"level"`//质量:0~标准
	Mode int `json:"mode"`//模式
	Wave string `json:"wave"`//频谱
	LrcZh string `json:"lrcZh"`//字幕~中文
	LrcEs string `json:"lrcEn"`//字母~英文
	Source int `json:"source"`//创作类型:0~原创,1~其它
}

var FILE ="SELECT id,content_id contentId,classify,cover,url,name,sufix,format,duration,width,height,size,rotate,bitrate,sample_rate sampleRate,level,mode,wave,lrc_zh lrcZh,lrc_es lrcEs,source FROM zd_file WHERE content_id=? AND classify=? order by create_time DESC LIMIT ? OFFSET ?"

var BLOG_UPDATE_FILE="UPDATE zd_blog SET url=?,name=?,sufix=?,format=?,duration=?,width=?,height=?,size=?,rotate=?,bitrate=?,sample_rate=?,level=?,mode=?,wave=?,lrc_zh=?,lrc_es=?,source=? WHERE id=?"


var file []File

func (this *File) TableName()string{
	return TableName("file")
}

func (this *File) Add()(int64,error){
	return orm.NewOrm().Insert(this)
}

func (this *File) Del()(int64,error){
	return orm.NewOrm().Delete(this)
}

func AddFile(filesJson string,id int64){
	beego.Info("filesJson:",filesJson)
	err := json.Unmarshal([]byte(filesJson),&file)
	beego.Info("err:",err)
	for k,v := range file	{
		if k == 0{
			beego.Info("--------jsd~v:",v)
			if len(v.LrcZh)==0 && len(v.Name) > 0{
				v.LrcZh = GetLrcStr(id,v.Name)
			}
			if _,err := SqlRaw(BLOG_UPDATE_FILE,[...]interface{}{v.Url,v.Name,v.Sufix,v.Format,v.Duration,v.Width,v.Height,v.Size,v.Rotate,v.Bitrate,v.SampleRate,v.Level,v.Mode,v.Wave,v.LrcZh,v.LrcEs,v.Source,id});err != nil{
				beego.Info("update~blogfile~err:",err)
			}
		}else{
			v.ContentId = id
			if _,err := v.Add();err != nil{
				beego.Info("add~file~err:",err)
			}
		}
		
	}
}

