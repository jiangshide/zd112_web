package controllers

type VideoController struct {
	BaseController
}

func (this *VideoController) Get() {
	this.nav()
	this.display("video/video")
}
