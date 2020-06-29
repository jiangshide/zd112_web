package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	"reflect"
	"github.com/jiangshide/GoComm/utils"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbHost := beego.AppConfig.String("db.host")
	dbPort := beego.AppConfig.String("db.port")
	dbUser := beego.AppConfig.String("db.user")
	dbPsw := beego.AppConfig.String("db.psw")
	dbName := beego.AppConfig.String("db.name")
	timeZone := beego.AppConfig.String("db.timezone")
	maxConn, _ := beego.AppConfig.Int("db.maxConn")
	maxIdle, _ := beego.AppConfig.Int("db.maxIdle")
	if dbPort == "" {
		dbPort = "3306"
	}
	dns := dbUser + ":" + dbPsw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	if timeZone != "" {
		dns += "&loc=" + url.QueryEscape(timeZone)
	}
	orm.RegisterDataBase("default", "mysql", dns, maxConn, maxIdle)
	orm.RegisterModel(new(User), new(UserLocation), new(UserProfile), new(Nav), new(Banner), new(Nation))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func Field(model interface{}) (fieldName string, fieldValue interface{}) {
	if model != nil {
		var field reflect.Type
		var value reflect.Value
		field = reflect.TypeOf(model).Elem()
		value = reflect.ValueOf(model).Elem()
		size := field.NumField()
		for i := 0; i < size; i++ {
			v := value.Field(i)
			fieldName = utils.StrFirstToLower(field.Field(i).Name)
			switch value.Field(i).Kind() {
			case reflect.Bool:
				fieldValue = v.Bool()
			case reflect.Int:
				if v.Int() != 0 {
					fieldValue = v.Int()
				}
			case reflect.Int64:
				if v.Int() != 0 {
					fieldValue = v.Int()
				}
			case reflect.String:
				fieldValue = v.String()
			}
			if fieldValue != nil && fieldValue != 0 {
				break
			}
		}
	}
	return
}

func SqlRaw(sql string,ids interface{})(res interface{},err error){
	res,err = orm.NewOrm().Raw(sql,ids).Exec()
	return
}


func SqlCount(table string,maps map[string]interface{})(count int64,err error){
	query := orm.NewOrm().QueryTable(table)
	for k,v := range maps{
		beego.Info("k:",k," | v:",v)
		query = query.Filter(k,v)
	}
	count,err = query.Count()
	return
}

func Sql(sql string,ids interface{})(res interface{},id int64,err error) {
	var maps []orm.Params
 	o := orm.NewOrm()
    id,err = o.Raw(sql,ids).Values(&maps)
   if len(maps) > 0{
   		res = maps[0]
   }
    return res,id,err
}

func SqlList(sql string,ids interface{}) (list *[]orm.Params,id int64,err error) {
 	var maps []orm.Params
 	o := orm.NewOrm()
    id,err = o.Raw(sql,ids).Values(&maps)
    return &maps,id,err
}
