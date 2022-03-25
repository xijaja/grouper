
<div align="center">
<a href=""><img alt="Stellar" src="https://axure.occo.pro/icon25.png" width="" /></a>
<br/>
<strong> 原型仔之友 </strong>
<h1>Grouper</h1>
</div>
<p align="center">
<a href=""><img alt="Build Status" src="https://circleci.com/gh/stellar/go.svg?style=shield" /></a>
<a href=""><img alt="Golang" src="https://img.shields.io/badge/Made%20with-Go-00ADD8.svg" /></a>
<a href=""><img alt="Maintained" src="https://img.shields.io/badge/Maintained%3F-yes-red.svg" /></a>
</p>

## 🥹 介绍
Grouper 是一个将本地某个特定的文件夹上传到云oss存储的服务，它会自动遍历文件夹下的文件内容，
作为一名产品经理(PM)，我使用 Axure 制作 PRD 并且它将导出 html 文件，但苦于难与团队共享。

天下产品人苦 Axure 久已，我先后尝试过：：Axure自带的云、国内的PmDaNiu、AxHub等，
他们要么是服务器在国外，要么是转而开始收费，或者只给你很小的空间、甚至限制你的文件数量。私有部署的话，费用更高。

当然，也有一些其他软件，他们是在线的，且支持团队协作。比如：墨刀、xiaopiu、MasterGo等原型设计共享工具；
但是，那样就只是产品原型了，如果你的产品原型和文档是分离的，选择他们自然是比较好的方式。
而我习惯是将文档、原型、注释、流程图、思维导图、外部链接等放在一份PRD中，我认为这将有利于程序猿们查看需求，
他们不需要为了某一个需求打开多份文档，反复比对；同时，如果变更，我和我的团队也将只变更一份文件足矣。

后来我买了台服务器，配置了 Nginx 和 SSL ，对于大多数产品人来说，这个方法的使用成本较高，你需要学习许多
与你工作不相关或者你也不感兴趣的内容。直到最近，我的服务器又要续费了……

转而，我开始使用各大云厂商的 OSS 静态文件托管服务。经过一番操作和体验，这东西基本上免费属于是，
但是这还不足够，所以"自己动手，丰衣足食"，我开发了这个软件，希以为原型仔之友，望与诸君共享。

## 🥳 优势
- **绑定域名**（支持自定义域名，实际上云厂商提供的）
- **无限空间**（OSS 静态储存服务至少 50G 起步）
- **私有部署**（你的文件永远是你的，整个仓库都是你的）
- **极速上传**（最高1024个并发同时上传，颤抖！懂？）
- **开源免费**（这个不用我解释了吧...2333......）
- **独家彩蛋**（互联网黑话弹窗，给我整不会了～）

## 😍 使用
第一步：下载并安装
1. 你可以选择 [点击这里](https://github.com/xiwuou/grouper/releases/tag/v1.0.0-beta) 下载你所需的程序。
2. 或者也可以自行编译程序，请继续阅读 编译。

第二步：注册云账户
1. 目前支持 [阿里云](https://www.aliyun.com) [腾讯云](https://cloud.tencent.com) [七牛云](https://www.qiniu.com) 
2. 注册完成拿到对应的字段资料，可以[参考教程]()

第三步：开始上传
1. 在 Grouper 的`设置`中选择你使用的服务，如`阿里云OSS`，并将配置参数填入。
2. 创建你的产品文档项目，名称中不要用中文。为了保持访问，项目名称不应该被修改。
3. 项目所在文件夹应选择你 Axure 导出 Html 静态文件的文件夹，被选中的文件夹中的所有内容将被上传。（ **特别注意**：不要填写你电脑的根目录或者仅填写 " / "）

备注：本程序的源代码主要仓库是 [Github](https://github.com/xiwuou/grouper) 
如果你受到网络影响也可以使用备份的 [Gitee](https://gitee.com/xiwuou/grouper) 

## 😴 编译
1. 如果你想自己编译这个程序，那么你需要：
```text
    go version >= 1.18
```
2. 克隆这个项目：
```text
    git clone https://github.com/xiwuou/grouper.git
```
3. 其次，我提供了两个不同的版本供你选择：
- GUI 图形用户界面程序
- CLI 命令行界面程序

### 编译GUI程序
a. 要编译带有GUI用户界面的程序，你需要进入`cd /grouper/cmd`目录中，执行命令：
```shell
    # for linux (由于我的电脑是mac所以没有修改，你可以以此创建mac或linux的程序)
    make default
    
    # for windows
    make windows
    
    # for all (这个命令将生成`.app`和`.exe`两个程序)
    make
```
b. 生成之后，将其拖入你的应用程序列表即可。如果是Windows系统，你可以直接打开。

### 编译CLI程序
a. 要编译带有GUI用户界面的程序，你需要进入`cd /grouper/cli`目录中，执行命令：
```shell
    # for linux
    go build -o grouper main.go
    
    # for windows
    go build -o grouper.exe main.go
```
b. 随后你将得到一个名为`grouper`的二进制可执行文件，要执行它，你可以使用`./grouper -h`获取帮助信息：
```text
  -n name   项目名称，请使用小写字母开头不含特殊符号，默认为文件夹名
  -p path   指定上传文件夹的路径，需为绝对路径，默认当前目录 (default ".")
  -v        显示出版本信息
  -version  显示出版本信息
```
c. 你可以将程序放置在你的`/usr/bin`或`/usr/local/bin`目录下，之后再次使用则不需要进入到程序所在目录
并且适应`./`前缀。Windows系统可能还需要设置环境变量。

## 🥴 开发者
如果你想要联系到我，交个朋友或者提个需求，抑或是提个BUG。左边是我的个人微信，右边是我的公众号。

<img src="https://axure.occo.pro/pm_xiwu.png" width="220" alt="我的个人微信"/>
<img src="https://axure.occo.pro/pmgouzi.jpg" width="229" alt="微信公众号：逆天思维产品汪"/>


## 🫡 感谢
🙏 UI支持 `github.com/AllenDang/giu` 

🙏 并发支持`github.com/panjf2000/ants`

🙏 阿里云SDK `github.com/aliyun/aliyun-oss-go-sdk`

🙏 七牛云SDK `github.com/qiniu/go-sdk`

🙏 腾讯云SDK `github.com/tencentyun/cos-go-sdk-v5`

