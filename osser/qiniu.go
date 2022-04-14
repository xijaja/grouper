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
func QiniuGetUpToken(qnCfg conf.QiniuOss) *qiniuOss {
	// 文件上传的上传策略
	putPolicy := storage.PutPolicy{
		Scope: qnCfg.BucketName, // 简单策略
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
}

// QiniuCoverUpload 七牛云上传
func QiniuCoverUpload(qnCfg conf.QiniuOss, couldFile, localFile string) {
	// 需要覆盖的文件名
	j := len("/Users/xiwu/Documents/Axure/MyDemo/")
	keyToOverwrite := localFile[j:]
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", qnCfg.BucketName, keyToOverwrite),
	}
	mac := qbox.NewMac(qnCfg.AccessKey, qnCfg.SecretKey)
	upToken := putPolicy.UploadToken(mac)
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
}
