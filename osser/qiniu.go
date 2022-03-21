package osser

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"grouper/conf"
)

type qiniuOss struct {
	upToken string
}

// QiniuGetUpToken 获取上传token
func QiniuGetUpToken() *qiniuOss {
	// 获取配置信息
	qnCfg := conf.Cfg.QiniuOss

	// 文件上传的上传策略
	putPolicy := storage.PutPolicy{
		Scope: qnCfg.BucketName,
	}
	mac := qbox.NewMac(qnCfg.AccessKey, qnCfg.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return &qiniuOss{upToken: upToken}
}

// QiniuGoUpload 上传
func (qn *qiniuOss) QiniuGoUpload(couldFile, localFile string) {
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
	err := formUploader.PutFile(context.Background(), &ret, qn.upToken, couldFile, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
