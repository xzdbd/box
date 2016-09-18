package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xzdbd/box/routers"
)

type loginInfo struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	beego.Run()
}
