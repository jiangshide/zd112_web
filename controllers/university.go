package controllers

type UniversityController struct {
	BaseController
}

func (this *UniversityController) Get() {
	this.nav()
	title1 := []string{"先秦", "汉朝", "隋朝", "唐朝", "宋朝", "金朝", "辽朝", "元朝", "清朝", "近当代", "辽朝"}
	this.Data["title1"] = title1

	title2 := []string{"诗经(305)", "屈原(27)", "荆轲(1)", "宋玉(1)", "李白(20)", "屈原(27)", "荆轲(1)", "宋玉(1)"}
	this.Data["title2"] = title2

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
	this.display("university/university")
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
