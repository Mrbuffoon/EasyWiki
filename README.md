# EasyWiki

Go语言版本的轻Wiki，可以将Markdown文件渲染成html并发布。通过配置webhooks可以实现自动拉取github/gitlab中内容进行发布。

NOTE：
> github的项目内必须要有一个RESUME.md文件，文件中列出要展示的目录列表（Markdown语法列表并链接），发布后将作为目录。

> 同时还需要一个README.md文件，可以简单描述一下这个项目的内容，发布后将作为主页

RESUME.md 示例：
```
# [go标准输入输出之占位符](go标准输入输出之占位符.md)

## [SDN初步调研](SDN初步调研.md)

### [代码管理中常用git操作](代码管理中常用git操作.md)

#### [多线程场景pymysql线程池使用的一点思考](多线程场景pymysql线程池使用的一点思考.md)

##### [go_sync](go_sync.md)

###### [groupcache_summury](groupcache_summury.md)

```

#### 依赖
* go 1.12 以上
* git安装并配置好，能够正常使用github
* nginx等能够提供Web服务的软件

#### 使用
* 配置 GOPATH 环境变量
* cd 进入 GOPATH/src 路径
* 下载代码

```
git clone git@github.com:Mrbuffoon/EasyWiki.git
```

* 安装

```
cd EasyWiki
make
make install
```

* 运行

首先自行设置配置文件/etc/easywiki/easywiki.ini
```
#EasyWiki默认数据路径
[SYSTEM]
MdPath = /var/easywiki/blogs

#默认日志输出路径
[LOG]
LogPath = /var/log/easywiki

#要同步的Git仓库地址
[GIT]
RepoAddr = git@github.com:Mrbuffoon/blogs.git

#Web网站根目录
[WEB]
WebRoot = /data/www

#监听端口，默认9090
[PORT]
Port = 9090
```

配置webhook( URL是 "ip:port/webhooks" )

最后运行即可
```
easywiki
```

