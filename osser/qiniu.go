package osser

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/rpc"
	"github.com/qiniu/go-sdk/v7/storage"
	"grouper/conf"
	"strings"
)

// ---------------------------------------------
// 上传
// ---------------------------------------------

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

// ---------------------------------------------
// 文件管理
// ---------------------------------------------

type qiniuOssBucketManager struct {
	BucketName    string
	BucketManager *storage.BucketManager
}

func QiniuBucketManager(qin conf.QiniuOss) *qiniuOssBucketManager {
	mac := qbox.NewMac(qin.AccessKey, qin.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return &qiniuOssBucketManager{
		BucketName:    qin.BucketName,
		BucketManager: bucketManager,
	}
}

// GetPrefixFiles 获取指定前缀的文件列表
func (manager *qiniuOssBucketManager) GetPrefixFiles(prefix string) (fileList []string, err error) {
	bucket := manager.BucketName
	limit := 1000
	// prefix := "qshell/"
	delimiter := ""
	// 初始列举marker为空
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := manager.BucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			return nil, err
			// break
		}
		// 对返回的文件列表进行处理
		for _, entry := range entries {
			// fmt.Println(entry.Key) // 打印条目
			fileList = append(fileList, entry.Key)
		}
		if hasNext {
			marker = nextMarker
		} else {
			// 列表结束，退出循环
			break
		}
	}
	return fileList, err
}

// GetPrefixFolders 获取带有指定前缀的文件夹列表
func (manager *qiniuOssBucketManager) GetPrefixFolders(prefix string) (folderList []string, err error) {
	files, filesErr := manager.GetPrefixFiles(prefix)
	for _, f := range files {
		folder := strings.Split(f, "/")
		fmt.Println(folder[0])
		folderList = append(folderList, folder[0])
	}
	return folderList, filesErr
}

// Delete 删除文件
func (manager *qiniuOssBucketManager) Delete(fileKeys []string) error {
	// 桶名
	bucket := manager.BucketName
	// 如果文件为空，则直接返回
	if fileKeys == nil {
		return nil
	}
	// 如果只有一个文件，则直接删除
	if len(fileKeys) == 1 {
		key := fileKeys[0]
		err := manager.BucketManager.Delete(bucket, key)
		if err != nil {
			return err
		}
	} else {
		// 如果有多个文件，则批量删除
		var ti int
		if len(fileKeys) > 1000 {
			ti = len(fileKeys)/900 + 1
			for i := 1; i < ti; i++ {
				start := (i - 1) * 900
				end := i*900 - 1
				keys := fileKeys[start:end]
				manager.executeDelete(keys)
			}
		} else {
			manager.executeDelete(fileKeys)
		}
	}
	return nil
}

// 执行删除
func (manager *qiniuOssBucketManager) executeDelete(keys []string) {
	// 桶名
	bucket := manager.BucketName
	// keys := fileKeys
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}
	rets, err := manager.BucketManager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*rpc.ErrorInfo); ok {
			for _, ret := range rets {
				// fmt.Printf("%d\n", ret.Code) // 200 为成功
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("批删除错误: %s", err)
		}
	}
}
