package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/xzdbd/box/controllers"
)

func init() {
	// validate if user is logged in.
	validateUserLogin()

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/disk", &controllers.DiskController{})
	beego.AutoRouter(&controllers.DiskController{})
}

func validateUserLogin() {
	var FilterUser = func(ctx *context.Context) {
		isLogin, ok := ctx.Input.Session("login").(string)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
		if ctx.Request.RequestURI == "/login" {
			if isLogin == "true" {
				ctx.Redirect(302, "/disk")
			}
		} else {
			if isLogin != "true" {
				ctx.Redirect(302, "/login")
			}
		}
	}

	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)
}
