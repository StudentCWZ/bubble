# bubble

## 仓库内容
- Go Web 基础教程 bubble 项目，提供如下功能：
	- 利用 Gin 框架 + MySQL 搭建一套通用的 Go Web 开发通用脚手架
	- 用户可以新建提醒清单，以及查询自己的提醒清单
	- bubble 项目具体代码详解可见 [Go Web 基础教程](https://www.bilibili.com/video/BV1gJ411p7xC?spm_id_from=333.999.0.0)，建议使用之前进行阅读

## 项目目录
- bubble 项目目录如下：
```
├── controller
│   └── controller.go
├── dao
│   └── mysql.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── todo.go
├── routers
│   └── routers.go
├── static
│   ├── css
│   │   ├── app.8eeeaf31.css
│   │   └── chunk-vendors.57db8905.css
│   ├── fonts
│   │   ├── element-icons.535877f5.woff
│   │   └── element-icons.732389de.ttf
│   └── js
│       ├── app.007f9690.js
│       └── chunk-vendors.ddcb6f91.js
└── templates
    ├── favicon.ico
    └── index.html
```

## 使用准备
- 克隆代码并进入 bubble 文件夹：
```
git clone git@github.com:StudentCWZ/bubble.git
cd bubble
```

## 使用要求

### 常规方式
- 这里使用常规方式来运行服务，常规方式要求有 Golang 环境、MySQL 环境，具体要求如下：
	- Golang >= 1.14
	- MySQL

### MySQL 安装配置
- 本地安装 MySQL、Docker 启动 MySQL、远程 MySQL 都是可以的，只要能正常连接使用即可，这里建议用 Docker 起一个 MySQL 。
- 需要在 MySQL 数据库中新建 bubble 库 。

### Golang 环境安装
- Golang 环境安装详细教程参考[相关博客](https://www.yuque.com/cuicuiaixiedaima/bhgmys/udybgl)

### 安装依赖包
- 要求：bubble 项目目录下
- 运行相关命令
```shell script
go mod tidy
```

### 运行服务
- 第一种：
```shell script
go run main.go
```

- 第二种：
```shell script
go build
bubble
```





