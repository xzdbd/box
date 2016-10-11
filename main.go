package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/postgres"
	_ "github.com/xzdbd/box/routers"
)

func main() {
	beego.Run()
}
