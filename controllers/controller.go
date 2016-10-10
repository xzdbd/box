package controllers

import (
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
	"github.com/xzdbd/box/models"
	"path"
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

	c.Data["ShareMessage"] = ""

	action := c.GetString("action")
	if action == "share" {
		objectName := c.GetString("objectName")
		fileName := path.Base(objectName)
		url, err := models.GetSharedUrl(objectName, fileName, 1)
		if err != nil {
			beego.Trace("共享失败：", err)
			c.Data["ShareMessage"] = renderShareFailMessage()
		}
		c.Data["ShareMessage"] = renderShareSuccessMessage(url)
	}

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
		beego.Trace("is folder?", isObjectFolder(object))
		if isObjectFolder(object) {
			htmlTpl += "<tr><td><a href=\"/disk/home/" + object.Key + "\"><i class=\"folder icon\">" + object.Key + "</i></a></td><td>" + object.LastModified.String() + "</td><td></td></tr>"
		} else {
			htmlTpl += "<tr><td><a href=\"/disk/home/" + object.Key + "\">" + object.Key + "</a></td><td>" + object.LastModified.String() + "</td><td><button class=\"ui primary button\" onclick=\"window.location.href='/disk/home/" + prefix + "?objectName=" + object.Key + "&action=share'\">共享</button></td></tr>"
		}

	}
	return htmlTpl
}

func renderShareSuccessMessage(url string) string {
	htmlTpl := "<div class=\"ui positive message\"><i class=\"close icon\"></i><div class=\"header\">共享完成!</div><p>请打开或复制以下链接来下载。</p><p><a href=" + url + ">" + url + "</a></p></div>"
	return htmlTpl
}

func renderShareFailMessage() string {
	htmlTpl := "<div class=\"ui positive message\"><i class=\"close icon\"></i><div class=\"header\">共享失败!</div></div>"
	return htmlTpl
}

func isObjectFolder(object minio.ObjectInfo) bool {
	if object.StorageClass == "STANDARD" {
		return false
	} else {
		return true
	}
}

func (c *MinioController) Get() {
	//models.ListBuckets()
	//models.PutObject()
	//models.ListObjects()
	//objects := models.GetUserObjects("bucket1", "", false)
	//beego.Trace(objects)
	url, _ := models.GetSharedUrl("test/logo", "logo.jpeg", 1)
	beego.Trace("url:", url)
	c.TplName = "login.tpl"
}
