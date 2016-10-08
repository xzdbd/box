package routers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/box/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/minio", &controllers.MinioController{})
}
