
# Go-Gin-Api
基于golang开源框架 gin封装的api框架

# 项目文档

[在线文章](https://juejin.cn/user/2313028196368445/posts)
## 1 项目简介

### 1.1 项目介绍

个人爱好，业务编码，喜欢的给个star
## 2 使用说明

```
- golang版本 >= v1.14
```
> 使用git 克隆

- 使用git克隆本项目
```git
    git clone https://github.com/18211167516/Go-Gin-Api.git
```
### 2.1 swagger自动化API文档

#### 2.1.1 安装 swagger
```bash
# 启用 Go Modules 功能
go env -w GO111MODULE=on 
# 配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.io,direct

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.1.2 生成API文档

```
swag init
```

> 执行上面的命令后，server目录下会出现docs文件夹，登录http://localhost:8080/swagger/index.html，即可查看swagger文档
## 3 技术选型

- 用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架
- 数据校验 `go-playground` gin框架内置，已支持错误信息转中文
- 数据库：采用`MySql`(5.6.44)版本，使用`gorm`实现对数据库的基本操作。
- API文档：使用`Swagger`构建自动化文档。
- 热重启`HotStart`,我自己写的一个
- 配置文件 使用`viper`解析
- 日志：使用`logrus`实现日志记录。

## 4. 项目架构
### 4.1 目录结构

```
    ├─app  	     （项目核心目录）
    |  ├─controller     （控制器）
    |  ├─middleware     （中间件）
    |  ├─models          （数据结构层）
    |  ├─request          （数据校验层）
    |  ├─services          （服务层）
    ├─config         （配置包）
    ├─core  	        （內核）
    ├─global            （全局变量）
    ├─initialize       (初始化)  
    ├─docs  	        （swagger文档目录）
    ├─routes         （路由）
    ├─tests          （测试）
    └─tool	        （公共功能）

```


