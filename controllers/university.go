package controllers

type UniversityController struct {
	BaseController
}

func (this *UniversityController) Get() {
	this.display("university/university")
}
