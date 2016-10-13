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
	beego.Router("/minio", &controllers.MinioController{})
}

func validateUserLogin() {
	var FilterUser = func(ctx *context.Context) {
		isLogin, ok := ctx.Input.Session("login").(string)
		beego.Trace("Login session info:", isLogin)
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
