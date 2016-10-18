package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Userinfo struct {
	Id           int64  `form:"-"`
	Username     string `form:"username"`
	Password     string `form:"password"`
	role         int64
	Lastmodified time.Time
	Created      time.Time
}

func (u *Userinfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *Userinfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}
