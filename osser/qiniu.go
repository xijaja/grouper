package osser

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func QiniuGetUpToken() (upToken string) {
	accessKey := "ZIj1s4Zj2i3djw8V90hI3zcL-YAe_9Ro2m43jdXq"
	secretKey := "BHY0QD2cejnoO_ACxGVsaT3db_aalrmAnv0gaQcJ"
	bucket := "myprd"

	// 文件上传的上传策略
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	return upToken
}

func QiniuGoUpload(upToken, couldFile, localFile string) {
	// 文件上传，资源管理等配置
	cfg := storage.Config{}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		// Params: map[string]string{
		// 	"x:name": "github logo",
		// },
	}

	// 开始上传
	err := formUploader.PutFile(context.Background(), &ret, upToken, couldFile, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
