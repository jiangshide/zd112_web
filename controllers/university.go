package controllers

import "github.com/astaxie/beego"

type UniversityController struct {
	BaseController
}

func (this *UniversityController) Get() {
	this.nav()
	this.filter()
	this.content()
	this.display("university/university")
}

type University struct {
	Id           int
	Name         string
	PoetId       int64
	PoetArrow    int
	Poet         [] Poet
	DynastyId    int64
	DynastyArrow int
	Dynasty      [] Dynasty
	AuthId       int64
	AuthArrow    int
	Auth         [] Auth
}

type Poet struct {
	Id    int
	Name  string
	Count int
}

type Dynasty struct {
	Id    int
	Name  string
	Count int
}

type Auth struct {
	Id    int
	Name  string
	Count int
}

func (this *UniversityController) filter() {
	var university University
	poetId := this.getInt64("poetId", 0)
	beego.Info("----------poetId:", poetId)
	university.PoetId = 4
	poetArr := []string{"年代诗人", "著名诗人", "诗词标签", "诗人故事"}
	for k, v := range poetArr {
		var poet Poet
		poet.Id = k
		poet.Name = v
		university.Poet = append(university.Poet, poet)
		if k > 10 {
			university.PoetArrow = 1
			break
		}
	}
	dynastyId := this.getInt64("dynastyId", 0)
	beego.Info("--------------dynastyId:", dynastyId)
	university.DynastyId = 12
	dynastyArr := []string{"先秦", "汉朝", "隋朝", "唐朝", "宋朝", "金朝", "辽朝", "元朝", "清朝", "近当代"}
	for k, v := range dynastyArr {
		var dynasty Dynasty
		dynasty.Id = k
		dynasty.Name = v
		university.Dynasty = append(university.Dynasty, dynasty)
		if k > 10 {
			university.DynastyArrow = 1
			break
		}
	}
	authId := this.getInt64("authId", 0)
	beego.Info("-----------authId:", authId)
	university.AuthId = 8
	authArr := []string{"诗经(305)", "屈原(27)", "荆轲(1)", "宋玉(1)", "李白(20)", "屈原(27)", "荆轲(1)", "宋玉(1)"}
	for k, v := range authArr {
		var auth Auth
		auth.Id = k
		auth.Name = v
		university.Auth = append(university.Auth, auth)
		if k > 10 {
			university.AuthArrow = 1
			break
		}
	}
	this.Data["university"] = university
}

func (this *UniversityController) content() {
	university := map[string]string{"/static/img/item1.png": "一剪梅", "/static/img/item2.png": "古来黎", "/static/img/item3.png": "李白", "/static/img/item4.png": "崧台"}
	content := make([][]*Content, 0)

	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	content = append(content, this.getItem(university))
	this.Data["universityContent"] = content
}

func (this *UniversityController) getItem(poetry map[string]string) []*Content {
	contentArr := make([] *Content, 0)
	for k, v := range poetry {
		content := new(Content)
		content.Id = 1
		content.Title = v
		content.Name = "李清照"
		content.Year = "宋代"
		content.Icon = k
		content.Like = "20万"
		content.Follow = "100万"
		content.Info = 80
		content.Before = "1天前"
		user := new(User)
		user.Id = 1
		user.Name = "春天来到"
		user.Icon = "/static/img/user_icon.png"
		content.User = user
		contentArr = append(contentArr, content)
	}
	return contentArr
}

type UniversityDetailController struct {
	BaseController
}

func (this *UniversityDetailController) Get() {
	this.nav()
	this.display("university/university_detail")
}
