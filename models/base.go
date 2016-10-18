package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	dbconn := beego.AppConfig.String("dbconnection")
	beego.Trace("Initialzing database. DB connection:", dbconn)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbconn)
	orm.RegisterModel(new(Userinfo))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}
