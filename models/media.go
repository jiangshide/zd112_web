package models

//http://gecimi.com/api/lyric/海阔天空

import(
	"github.com/astaxie/beego"
)

type LrcCode struct{
	Count int `json:"count"`
	Code int `json:"code"`
	Result 	[]Lrc `json:"result"`
}

type Lrc struct{
	Aid int64 `json:"aid"`
	ArtistId int `json:"artist_id"`
	Lrc string 	`json:"lrc"`
	Sid int64 `json:"Sid"`
	Song string `json:"song"`		
}

var UPDATE_LRC = "UPDATE zd_blog SET name=?,lrc_zh=?,lrc_es=? WHERE id=?"
var UPDATE_LRC_ZH = "UPDATE zd_blog SET name=?,lrc_zh=? WHERE id=?"
var UPDATE_LRC_ES = "UPDATE zd_blog SET name=?,lrc_zh=? WHERE id=?"

func GetLrc(id int64,name string,lrcCode *LrcCode){
	http := new(Http)
	http.Url = "http://gecimi.com/api/lyric/"+name
	http.Model = lrcCode
	if err := http.Get();err != nil{
		beego.Info("err:",err)
	}else{
		if id > 0 && len(lrcCode.Result) > 0 {
			SqlRaw(UPDATE_LRC_ZH,[...]interface{}{lrcCode.Result[0].Song,lrcCode.Result[0].Lrc,id})
		}
	}
}

func GetLrcStr(id int64,name string)string{
	var lrcCode LrcCode
	lrc := ""
	http := new(Http)
	http.Url = "http://gecimi.com/api/lyric/"+name
	http.Model = &lrcCode
	if err := http.Get();err != nil{
		beego.Info("err:",err)
	}else{
		if id > 0 && len(lrcCode.Result) > 0 {
			lrc = lrcCode.Result[0].Lrc
		}
	}
	return lrc
}