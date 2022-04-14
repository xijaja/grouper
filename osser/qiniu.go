package osser

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"grouper/conf"
)

type qiniuOss struct {
	qnCfg *conf.QiniuOss
}

// QiniuGetUpToken 获取上传token
func QiniuGetUpToken(qin conf.QiniuOss) *qiniuOss {
	return &qiniuOss{
		qnCfg: &qin,
	}
}

// QiniuCoverUpload 七牛云上传
func (qn *qiniuOss) QiniuCoverUpload(couldFile, localFile string) {
	// 需要覆盖的文件名
	keyToOverwrite := couldFile
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", qn.qnCfg.BucketName, keyToOverwrite),
	}
	mac := qbox.NewMac(qn.qnCfg.AccessKey, qn.qnCfg.SecretKey)
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
