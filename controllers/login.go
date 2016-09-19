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
		c.Redirect("http://23.88.238.182:9000", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
	//c.Data["loginStatus"] = loginStatus
	//c.TplName = "portal.tpl"
}

func (c *MinioController) Get() {
	models.ListBuckets()
	models.ListObjects()
	c.TplName = "login.tpl"
}
