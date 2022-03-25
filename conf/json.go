package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

// ---------------------------------------------
// 初始化
// ---------------------------------------------

var jsonFile string // 配置文件路径
var DataInfo *Data  // 配置信息

func init() {
	// 获取用户主目录 u.HomeDir
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile = u.HomeDir + "/.grouper.json"
	DataInfo = ReadData() // 初始化配置信息
}

// ---------------------------------------------
// 配置参数结构体
// ---------------------------------------------

// Project 项目
type Project struct {
	Name      string `json:"name"`       // 项目名称
	UpType    string `json:"up_type"`    // 上传服务类型，必须为：阿里云OSS、腾讯云COS、七牛云OSS，中的一个
	LocalFile string `json:"local_file"` // 本地地址
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

// ---------------------------------------------
// 读取配置
// ---------------------------------------------

// ReadData 读取配置
func ReadData() *Data {
	// 判断是否存在配置文件
	_, err := os.Stat(jsonFile)
	// 如果不存在则写入后再读取
	if os.IsNotExist(err) {
		dataBytes := `{
			"projects": [
				{
					"name": "test1",
					"up_type": "阿里云OSS",
					"local_file": "/您的主目录/您的子目录/同名文件夹test1"
				},
				{
					"name": "test2",
					"up_type": "阿里云OSS",
					"local_file": "/您的主目录/您的子目录/同名文件夹test2"
				}
			],
			"up_service": {
				"tencent_cos": {
					"bucket_name": "",
					"cos_region": "",
					"secret_id": "",
					"secret_key": "",
					"domain": ""
				},
				"aliyun_oss": {
					"endpoint": "",
					"key_id": "",
					"key_secret": "",
					"bucket_name": "",
					"domain": ""
				},
				"qiniu_oss": {
					"access_key": "",
					"secret_key": "",
					"bucket_name": "",
					"domain": ""
				}
			}
		}`
		// fmt.Println("dataBytes:", dataBytes)
		_ = ioutil.WriteFile(jsonFile, []byte(dataBytes), 0666)
	}
	// 否则直接读取json文件
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal("打开配置文件报错：", err)
	}
	// 绑定到结构体
	var info Data
	err = json.Unmarshal(jsonData, &info)
	if err != nil {
		log.Fatal("数据错误：", err)
	}
	return &info
}

// ---------------------------------------------
// 写入及更新配置
// ---------------------------------------------

// UpdateAliyunOss 更新阿里云OSS配置
func (ali *AliyunOss) UpdateAliyunOss() {
	DataInfo.UpService.AliyunOss = *ali
	resetJsonFile()
}

// UpdateTencentCos 更新腾讯云COS配置
func (ten *TencentCos) UpdateTencentCos() {
	DataInfo.UpService.TencentCos = *ten
	resetJsonFile()
}

// UpdateQiniuOss 更新七牛云OSS配置
func (qin *QiniuOss) UpdateQiniuOss() {
	DataInfo.UpService.QiniuOss = *qin
	resetJsonFile()
}

// AddOneProject 添加一个项目
func (p *Project) AddOneProject() {
	if p.UpType == "" {
		p.UpType = "阿里云OSS" // 默认选择
	}
	DataInfo.Projects = append(DataInfo.Projects, *p)
	resetJsonFile()
}

// UpdateOneProject 更新一个项目
func (p *Project) UpdateOneProject() {
	for i := 0; i < len(DataInfo.Projects); i++ {
		if DataInfo.Projects[i].Name == p.Name {
			DataInfo.Projects[i] = *p
		}
	}
	resetJsonFile()
}

// DeleteOneProject 删除一个项目
func (p *Project) DeleteOneProject() {
	var num int
	for i := 0; i < len(DataInfo.Projects); i++ {
		if DataInfo.Projects[i].Name == p.Name {
			num = i
			fmt.Println(p.Name)
		}
	}
	DataInfo.Projects = append(DataInfo.Projects[:num], DataInfo.Projects[num+1:]...)
	fmt.Println(DataInfo.Projects)
	resetJsonFile()
}

// 重置配置文件
func resetJsonFile() {
	data, err := json.MarshalIndent(DataInfo, "", "	")
	if err != nil {
		fmt.Println("错误：", err)
	}
	err = ioutil.WriteFile(jsonFile, data, 0777)
	if err != nil {
		fmt.Println("错误：", err)
	}
}
