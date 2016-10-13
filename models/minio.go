package models

import (
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
	"net/url"
	"os"
	"path"
	"time"
)

var minioClient *minio.Client
var err error

func init() {
	endpoint := "23.88.238.182:9000"
	accessKeyID := "IS9ICPLHJPDUJH9CZ7WS"
	secretAccessKey := "Ku1G/oUb+0BFjRRenNzjHOu+xtGG9Z/ZK3uXSGM8"
	useSSL := false

	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		beego.Trace("初始化minio client失败：", err.Error())
	}
	beego.Trace("初始化minio client完成：", minioClient)
}

func GetUserObjects(bucketName string, objectPrefix string, isRecursive bool) []minio.ObjectInfo {
	doneCh := make(chan struct{})
	var objects []minio.ObjectInfo

	defer close(doneCh)

	objectCh := minioClient.ListObjects("bucket1", objectPrefix, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			beego.Trace("ListObjects Error:", object.Err)
			continue
		}
		objects = append(objects, object)
	}
	return objects
}

func GetSharedUrl(objectName string, fileName string, expiryDays int) (string, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

	// Generates a presigned url which expires in 30 days.
	presignedURL, err := minioClient.PresignedGetObject("bucket1", objectName, time.Second*24*60*60*time.Duration(expiryDays), reqParams)
	if err != nil {
		beego.Error("生成共享url出错：", err.Error())
		return "", err
	}
	return presignedURL.String(), nil
}

//for test
func ListBuckets() {
	beego.Trace("开始查询buckets...")
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		beego.Trace("ListBuckets Error:", err.Error())
	}
	for _, bucket := range buckets {
		beego.Trace("buckets查询结果:", bucket)
	}
}

//for test
func ListObjects() {
	beego.Trace("开始查询objects...")

	doneCh := make(chan struct{})

	defer close(doneCh)

	isRecursive := false
	objectCh := minioClient.ListObjects("bucket1", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			beego.Trace("ListObjects Error:", object.Err)
			return
		}
		beego.Trace("objects查询结果:", object.Key, object.StorageClass, object.LastModified)
		//beego.Trace("文件路径：", filepath.Dir(object.Key), filepath.Base(object.Key))
		beego.Trace("文件路径：", path.Dir(object.Key), path.Base(object.Key))
	}
}

//for test
func PutObject() {
	beego.Trace("开始上传objects...")
	file, err := os.Open("static/img/logo.jpeg")
	if err != nil {
		beego.Trace("打开本地文件错误", err.Error())
		return
	}
	defer file.Close()

	n, err := minioClient.PutObject("bucket1", "test/logo2", file, "application/octet-stream")
	if err != nil {
		beego.Trace("上传错误", err.Error())
		return
	}
	beego.Trace("上传objects完成", n)
}
