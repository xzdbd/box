package models

import (
	"github.com/astaxie/beego"
	"github.com/minio/minio-go"
	"net/url"
	"os"
	"path"
	"time"
)

type MinioObject struct {
	minio.ObjectInfo
	IsFolder bool
}

type FileObject struct {
	MinioObject
	Dir  string
	Name string
}

type FolderObject struct {
	MinioObject
	Name string
}

type MyObjects struct {
	FolderObjects []FolderObject
	FileObjects   []FileObject
}

var minioClient *minio.Client

func init() {
	endpoint := beego.AppConfig.String("minioserver")
	accessKeyID := beego.AppConfig.String("minioaccesskey")
	secretAccessKey := beego.AppConfig.String("miniosecretkey")
	useSSL, er := beego.AppConfig.Bool("miniossl")
	if er != nil {
		beego.Warning("get AppConfig: miniossl failed, set to false by default. Error:", er)
		useSSL = false
	}

	var err error
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		beego.Error("Initialize Minio client failed. Error:", err)
	}
	beego.Trace("Initialize Minio client completed. Minio client info:", minioClient)
}

func (m *MyObjects) GetObjects(bucketName string, objectPrefix string, isRecursive bool) {
	doneCh := make(chan struct{})

	defer close(doneCh)

	objectCh := minioClient.ListObjects("bucket1", objectPrefix, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			beego.Error("ListObjects Error:", object.Err)
			continue
		}
		minioObject := new(MinioObject)
		minioObject.ObjectInfo = object
		minioObject.SetIsFolder()

		if minioObject.IsFolder {
			folderObject := new(FolderObject)
			folderObject.MinioObject = *minioObject
			folderObject.SetName()
			m.FolderObjects = append(m.FolderObjects, *folderObject)
		} else {
			fileObject := new(FileObject)
			fileObject.MinioObject = *minioObject
			fileObject.SetName()
			fileObject.SetDir()
			m.FileObjects = append(m.FileObjects, *fileObject)
		}
	}
}

func (m MyObjects) RenderMyObjects() string {
	htmlTpl := ""
	for i := 0; i < len(m.FolderObjects); i++ {
		folderObject := m.FolderObjects[i]
		htmlTpl += "<tr><td><a href=\"/disk/home/" + folderObject.Key + "\"><i class=\"folder icon\">" + folderObject.Name + "</i></a></td><td>" + folderObject.LastModified.String() + "</td><td></td></tr>"
	}
	for i := 0; i < len(m.FileObjects); i++ {
		fileObject := m.FileObjects[i]
		htmlTpl += "<tr><td>" + fileObject.Name + "</td><td>" + fileObject.LastModified.String() + "</td><td><button class=\"ui primary button\" onclick=\"window.location.href='/disk/home/" + fileObject.Dir + "?objectName=" + fileObject.Key + "&action=share'\">共享</button></td></tr>"
	}
	return htmlTpl
}

func (m *MinioObject) SetIsFolder() {
	if m.StorageClass == "STANDARD" {
		m.IsFolder = false
	} else {
		m.IsFolder = true
	}
}

func (f *FileObject) SetName() {
	f.Name = path.Base(f.Key)
}

func (f *FolderObject) SetName() {
	f.Name = path.Base(f.Key)
}

func (f *FileObject) SetDir() {
	f.Dir = path.Dir(f.Key)
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
