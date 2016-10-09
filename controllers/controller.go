package controllers

import (
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
	"github.com/xzdbd/box/models"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

type MinioController struct {
	beego.Controller
}

type DiskController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	u := models.Userinfo{}
	if err := c.ParseForm(&u); err != nil {
		beego.Trace(err.Error())
	}

	loginStatus := models.ValidateUserLogin(u)

	if loginStatus {
		c.Redirect("/disk", 302)
	} else {
		c.Data["isLoginFail"] = true
	}
	c.TplName = "login.tpl"
}

func (c *DiskController) Get() {
	c.Redirect("/disk/home", 302)
}

func (c *DiskController) Home() {
	params := c.Ctx.Input.Params()
	beego.Trace("params:", params)
	prefix := ""
	for i := 0; ; i++ {
		v, ok := params[strconv.Itoa(i)]
		if ok {
			prefix += v + "/"
		} else {
			break
		}
	}
	beego.Trace("prefix:", prefix)
	objects := models.GetUserObjects("bucket1", prefix, false)
	objectHtmlTpl := renderUserObjects(objects, prefix)
	c.Data["UserObjects"] = objectHtmlTpl
	c.TplName = "home.tpl"
}

func renderUserObjects(objects []minio.ObjectInfo, prefix string) string {
	htmlTpl := ""
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		beego.Trace(object.StorageClass, object.Key, object.Owner)
		htmlTpl += "<tr><td><a href=\"/disk/home/" + prefix + object.Key + "\"><i class=\"folder icon\">" + object.Key + "</i></a></td><td>" + object.LastModified.String() + "</td><td><button class=\"ui primary button\">共享</button></td></tr>"
	}
	return htmlTpl
}

func isObjectFolder(object minio.ObjectInfo) bool {
	if object.StorageClass == "STANDARD" {
		return true
	} else {
		return false
	}
}

func (c *MinioController) Get() {
	//models.ListBuckets()
	//models.PutObject()
	//models.ListObjects()
	//objects := models.GetUserObjects("bucket1", "", false)
	//beego.Trace(objects)
	url := models.GetSharedUrl("test/logo", "logo.jpeg", 1)
	beego.Trace("url:", url)
	c.TplName = "login.tpl"
}
