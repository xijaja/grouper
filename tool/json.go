package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Project 项目
type Project struct {
	Name      string `json:"name"`       // 项目名称
	UpType    string `json:"up_type"`    // 上传服务类型，必须为：阿里云OSS、腾讯云COS、七牛云OSS，中的一个
	LocalFile string `json:"local_file"` // 本地地址
	VisitAddr string `json:"visit_addr"` // 查看地址
}

// TencentCos 腾讯云对象储存cos配置
type TencentCos struct {
	BucketName string `json:"bucket_name"` // 桶名
	CosRegion  string `json:"cos_region"`  // 区域
	SecretID   string `json:"secret_id"`   // id
	SecretKey  string `json:"secret_key"`  // key
	Domain     string `json:"domain"`      // 域名地址
}

// AliyunOss 阿里云对象储存oss配置
type AliyunOss struct {
	Endpoint   string `json:"endpoint"`    // 地域节点地址
	KeyID      string `json:"key_id"`      // oss的key
	KeySecret  string `json:"key_secret"`  // oss的secret
	BucketName string `json:"bucket_name"` // 存储桶的名字
	Domain     string `json:"domain"`      // 域名地址
}

// QiniuOss 七牛云对象储存oss配置
type QiniuOss struct {
	AccessKey  string `json:"access_key"`  // ak
	SecretKey  string `json:"secret_key"`  // sk
	BucketName string `json:"bucket_name"` // 桶名
	Domain     string `json:"domain"`      // 域名地址
}

// Data Json配置信息
type Data struct {
	Projects  []Project `json:"projects"`
	UpService struct {
		TencentCos TencentCos `json:"tencent_cos"` // 腾讯云对象储存cos配置
		AliyunOss  AliyunOss  `json:"aliyun_oss"`  // 阿里云对象储存oss配置
		QiniuOss   QiniuOss   `json:"qiniu_oss"`   // 七牛云对象储存oss配置（还没有完成，暂时不可用）
	} `json:"up_service"`
}

// ReadData 读取配置
func ReadData() *Data {
	// 读取json文件 todo 根据系统不同，如果没有则自动创建至特定目录
	jsonData, err := ioutil.ReadFile("app/grouper.json")
	if err != nil {
		fmt.Println("打开配置文件报错：", err)
	}
	// 绑定到结构体
	var info Data
	err = json.Unmarshal(jsonData, &info)
	return &info
}

// WriteData todo 根据类型写入配置
func WriteData(a any) {
	switch a.(type) {
	case AliyunOss:
	case TencentCos:
	case QiniuOss:
	default:
		fmt.Println("传入参数有误")
	}
	// var dat Data
	// // append(dat.Projects, )
	// err := ioutil.WriteFile("app/grouper.json", "", os.ModeAppend)
}
