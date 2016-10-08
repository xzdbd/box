package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/box/models"
)

type LoginController struct {
	beego.Controller
}

type MinioController struct {
	beego.Controller
}

type HomeController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	u := models.Userinfo{}
	if err := c.ParseForm(&u); err != nil {
		beego.Trace(err.Error())
	}

	loginStatus := models.ValidateUserLogin(u)

	if loginStatus {
		c.Redirect("/home", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
}

func (c *HomeController) Get() {
	c.TplName = "home.tpl"
}

func (c *MinioController) Get() {
	//models.ListBuckets()
	//models.PutObject()
	//models.ListObjects()
	objects := models.GetUserObjects("bucket1", "", false)
	beego.Trace(objects)
	c.TplName = "login.tpl"
}
