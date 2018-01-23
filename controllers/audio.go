package controllers

type AudioController struct{
	BaseController
}

func (this *AudioController) Get(){
	this.nav()
	this.display("audio/audio")
}