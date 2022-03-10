package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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
	AliyunOss struct {
		Endpoint   string `yaml:"endpoint"`
		KeyID      string `yaml:"key_id"`
		KeySecret  string `yaml:"key_secret"`
		BucketName string `yaml:"bucket_name"`
		VisitAddr  string `yaml:"visit_addr"`
	} `yaml:"aliyun_oss"`
	QiniuOss struct {
		AccessKey  string `yaml:"access_key"`
		SecretKey  string `yaml:"secret_key"`
		BucketName string `yaml:"bucket_name"`
	} `yaml:"qiniu_oss"`
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
	address := Cfg.AliyunOss.VisitAddr
	if address[len(address)-1:] != "/" {
		return address + "/" + name
	} else {
		return address + name
	}
}
