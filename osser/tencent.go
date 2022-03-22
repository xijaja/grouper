package osser

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"grouper/conf"
	"net/http"
	"net/url"
)

type tencentCos struct {
	client *cos.Client
}

// CosClient 获取cos句柄
func CosClient(txCos conf.TencentCos) (tx *tencentCos) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为用户真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(fmt.Sprintf("https://%v.cos.%v.myqcloud.com", txCos.BucketName, txCos.CosRegion))
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://service.cos.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  txCos.SecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: txCos.SecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	return &tencentCos{client: client}
}

// Upload 上传
func (tx *tencentCos) Upload(key, file string) (ok bool) {
	// 上传配置
	opt := &cos.MultiUploadOptions{
		ThreadPoolSize: 1024,
	}
	// 上传
	_, _, err := tx.client.Object.Upload(
		context.Background(), key, file, opt,
	)
	// 处理错误
	if err != nil {
		panic(err)
	}
	return true
}
