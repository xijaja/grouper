package osser

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"grouper/common/conf"
	"net/http"
	"net/url"
)

// ---------------------------------------------
// 上传
// ---------------------------------------------

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

// ---------------------------------------------
// 文件管理
// ---------------------------------------------

// ObjList 获取腾讯云 cos 文件列表
func (tx *tencentCos) ObjList() (list []TencentCosFile) {
	opt := &cos.BucketGetOptions{
		Prefix:    "",
		Delimiter: "/",
		MaxKeys:   100,
	}
	result, _, err := tx.client.Bucket.Get(context.Background(), opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, content := range result.Contents {
		// fmt.Printf("文件: %v\n", content.Key)
		list = append(list, TencentCosFile{
			Object: content.Key,
			Types:  0,
		})
	}
	// common prefix表示表示被delimiter截断的路径, common prefix则表示所有子目录的路径
	for _, commonPrefix := range result.CommonPrefixes {
		// fmt.Printf("项目: %v\n", commonPrefix)
		list = append(list, TencentCosFile{
			Object: commonPrefix,
			Types:  1,
		})
	}
	return list
}

// TencentCosFile 腾讯云 cos 文件列表
type TencentCosFile struct {
	Object string // 文件名
	Types  int    // 0: 文件 1: 项目
}

// ObjDelete 删除腾讯云 cos 文件
func (tx *tencentCos) ObjDelete(key string) (ok bool) {
	// 查询对象是否存在
	ok, err := tx.client.Object.IsExist(context.Background(), key)
	if err != nil {
		fmt.Println("删除文件失败：", err)
		return false
	}
	if !ok {
		fmt.Println("文件或项目不存在：", key)
		return false
	}

	// 开始删除
	fmt.Printf("正在删除腾讯云 cos 上的 %v ，请稍候～\n", key)
	if key[len(key)-1] == '/' {
		// 表示是文件夹
		var marker string
		opt := &cos.BucketGetOptions{
			Prefix:  key,
			MaxKeys: 1000,
		}
		isTruncated := true
		for isTruncated {
			opt.Marker = marker
			v, _, err := tx.client.Bucket.Get(context.Background(), opt)
			if err != nil {
				fmt.Println(err)
				break
			}
			for _, content := range v.Contents {
				_, err = tx.client.Object.Delete(context.Background(), content.Key)
				if err != nil {
					fmt.Println(err)
				}
			}
			isTruncated = v.IsTruncated
			marker = v.NextMarker
		}
	} else {
		// 否则是文件
		_, err := tx.client.Object.Delete(context.Background(), key)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
