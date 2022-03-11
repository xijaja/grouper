package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Cfg 声明配置
var Cfg *MyConfig

// 初始化配置
func init() {
	c := &MyConfig{}
	Cfg = c.getMyConfig()
}

// MyConfig 配置文件结构体
type MyConfig struct {
	RootVisit  string `yaml:"root_visit"`
	TencentCos struct {
		BucketName string `yaml:"bucket_name"` // 桶名
		CosRegion  string `yaml:"cos_region"`  // 区域
		SecretID   string `yaml:"secret_id"`   // id
		SecretKey  string `yaml:"secret_key"`  // key
	} `yaml:"tencent_cos"` // 腾讯云对象储存cos配置
	AliyunOss struct {
		Endpoint   string `yaml:"endpoint"`    // 地域节点地址
		KeyID      string `yaml:"key_id"`      // oss的key
		KeySecret  string `yaml:"key_secret"`  // oss的secret
		BucketName string `yaml:"bucket_name"` // 存储桶的名字
	} `yaml:"aliyun_oss"` // 阿里云对象储存oss配置
	QiniuOss struct {
		AccessKey  string `yaml:"access_key"`  // ak
		SecretKey  string `yaml:"secret_key"`  // sk
		BucketName string `yaml:"bucket_name"` // 桶名
	} `yaml:"qiniu_oss"` // 七牛云对象储存oss配置（还没有完成，暂时不可用）
}

// 读取配置并绑定结构体
func (m *MyConfig) getMyConfig() *MyConfig {
	// 读取yaml文件到缓存
	yamlFile, err := ioutil.ReadFile("uper.yaml")
	if err != nil {
		fmt.Println("没有找到配置文件：", err)
	}
	// yaml文件内容映射到结构体中
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println("绑定配置参数错误：", err.Error())
	}
	return m
}

// Addr 访问地址
func Addr(name string) (addr string) {
	address := Cfg.RootVisit
	if address == "" {
		return name
	}
	if address[len(address)-1:] != "/" {
		return address + "/" + name
	} else {
		return address + name
	}
}

// Version 版本信息
func Version() {
	fmt.Println("版本号：v1.0-20210311")
	fmt.Println("开发者：習武（公众号：逆天思维产品汪）")
	fmt.Println("使用说明：xxx.xxx")
	fmt.Println("Github地址：https://github.com/xiwuou/uper")
	fmt.Println("感谢Star 🌟  欢迎Fork 👏")
	// 退出程序
	os.Exit(0)
}
