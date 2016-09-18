package test

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/box/models"
	_ "github.com/xzdbd/box/routers"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestListbucket(t *testing.T) {
	models.ListBuckets()
}
