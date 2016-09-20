package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/xzdbd/box/utils"
)

type Userinfo struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func init() {
	dbconn := beego.AppConfig.String("dbconnection")
	beego.Trace("获取DB连接信息:", dbconn)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dbconn)
	orm.RegisterModel(new(Userinfo))
	//orm.Debug = true
}

func ValidateUserLogin(u Userinfo) bool {
	u.Password = utils.EncryptPassword(u.Password)
	beego.Trace("获取加密后的密码:", u.Password)

	beego.Trace("正在验证用户名密码")
	o := orm.NewOrm()
	err := o.Read(&u, "Username", "Password")

	if err == orm.ErrNoRows {
		beego.Trace("验证用户名密码结果：未查询到用户")
	} else if err == orm.ErrMissPK {
		beego.Trace("验证用户名密码结果：找不到主键")
	} else {
		if u.Id > 0 {
			beego.Trace("验证用户名密码结果：验证成功", "ID：", u.Id, "用户名：", u.Username)
			return true
		}
	}
	return false
}
