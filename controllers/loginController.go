package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/box/models"
	"github.com/xzdbd/box/utils"
)

type LoginController struct {
	beego.Controller
}

type DiskController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	user := models.Userinfo{}
	if err := c.ParseForm(&user); err != nil {
		beego.Error(err.Error())
	}

	//loginStatus := models.ValidateUserLogin(user)
	loginStatus := validateUserLogin(&user)

	if loginStatus {
		beego.Trace("登陆成功！设置session为true")
		c.SetSession("login", "true")
		c.Redirect("/disk", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
}

func validateUserLogin(u *models.Userinfo) bool {
	u.Password = utils.EncryptPassword(u.Password)
	if err := u.Read("Username", "Password"); err != nil {
		beego.Trace("User validata failed:", err)
		return false
	}
	beego.Trace("User login success, user:", u.Username)
	return true
}
