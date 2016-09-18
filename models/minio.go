package models

import (
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
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

func ListObjects() {
	beego.Trace("开始查询objects...")

	doneCh := make(chan struct{})

	defer close(doneCh)

	isRecursive := true
	objectCh := minioClient.ListObjects("bucket1", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			beego.Trace("ListObjects Error:", object.Err)
			return
		}
		beego.Trace("objects查询结果:", object)
	}
}
