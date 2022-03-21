package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Data struct {
	Projects []struct {
		ProjectId int    `json:"project_id"` // 项目ID
		Name      string `json:"name"`       // 项目名称
		UpType    string `json:"up_type"`    // 上传服务类型：必须为 tencent alioss qiniu 中的一个
		LocalFile string `json:"local_file"` // 本地地址
		VisitAddr string `json:"visit_addr"` // 查看地址
	} `json:"projects"`
	UpService struct {
		TencentCos struct {
			BucketName string `json:"bucket_name"` // 桶名
			CosRegion  string `json:"cos_region"`  // 区域
			SecretID   string `json:"secret_id"`   // id
			SecretKey  string `json:"secret_key"`  // key
			Domain     string `json:"domain"`      // 域名地址
		} `json:"tencent_cos"` // 腾讯云对象储存cos配置
		AliyunOss struct {
			Endpoint   string `json:"endpoint"`    // 地域节点地址
			KeyID      string `json:"key_id"`      // oss的key
			KeySecret  string `json:"key_secret"`  // oss的secret
			BucketName string `json:"bucket_name"` // 存储桶的名字
			Domain     string `json:"domain"`      // 域名地址
		} `json:"aliyun_oss"` // 阿里云对象储存oss配置
		QiniuOss struct {
			AccessKey  string `json:"access_key"`  // ak
			SecretKey  string `json:"secret_key"`  // sk
			BucketName string `json:"bucket_name"` // 桶名
			Domain     string `json:"domain"`      // 域名地址
		} `json:"qiniu_oss"` // 七牛云对象储存oss配置（还没有完成，暂时不可用）
	} `json:"up_service"`
}

// ReadData 读取配置
func ReadData() *Data {
	// 读取json文件
	jsonData, err := ioutil.ReadFile("app/grouper.json")
	if err != nil {
		fmt.Println("打开配置文件报错：", err)
	}
	// 绑定到结构体
	var info Data
	err = json.Unmarshal(jsonData, &info)
	return &info
}

// WriteData todo 写入配置
func WriteData() {
	// var dat Data
	// // append(dat.Projects, )
	// err := ioutil.WriteFile("app/grouper.json", "", os.ModeAppend)
}
