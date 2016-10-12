package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/postgres"
	_ "github.com/xzdbd/box/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/server.log"}`)
	beego.BeeLogger.Async()
	beego.Run()
}
