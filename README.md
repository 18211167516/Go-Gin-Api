
# Go-Gin-Api
基于golang开源框架 gin封装的api框架

# 项目文档

[在线文章](https://juejin.cn/user/2313028196368445/posts)
## 1 项目简介

### 1.1 项目介绍

个人爱好，业务编码，喜欢的给个star
## 2 使用说明

```
- golang版本 >= v1.16
```
1. 下载

- 使用git克隆本项目
```git
    git clone https://github.com/18211167516/Go-Gin-Api.git
```
2.  执行sql文件

```
static/config/init.sql
```

3. 登录后台

```
账号：admin
密码：Admin123
```
4. 贡献指南
   
 
 * issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。
 * pull Request  请先 fork 一份到自己的项目下，不要直接在仓库下建分支。
 * commit 信息要以[文件名]: 描述信息 的形式填写，例如 README.md: fix xxx bug。
 * 确保 PR 是提交到 dev 分支，而不是 master 分支。
 * 如果是修复 bug，请在 PR 中给出描述信息。

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

## 2.2 热编译

```
go get github.com/pilu/fresh

进入项目执行

fresh

配置文件 是runner.conf
```
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
    |  ├─models         （数据结构层）
    |  ├─request        （数据请求层，定义特殊请求结构体以及数据校验）
    |  ├─request        （数据展示层定义结构体）
    |  ├─services       （服务层）
    ├─config            （配置包）
    ├─core  	        （內核）
    ├─docs  	        （swagger文档目录）
    ├─global            （全局变量）
    ├─initialize         (初始化)  
    ├─routes            （路由）
    ├─static            （静态文件包括config目录）
    ├─templates         （模板）
    ├─tests             （测试）
    └─tool	            （工具）

```
### 4.2 组件支持

1. 支持[Swagger](https://github.com/swaggo/gin-swagger)接口文档生成
2. 支持[jwt](https://github.com/golang-jwt/jwt)鉴权
3. 支持[logrus](https://github.com/sirupsen/logrus)(准备替换成zap)
4. 支持[viper](https://github.com/spf13/viper) 配置文件解析
5. 支持go1.6.0 go:embed特性,打包包含静态文件
6. 支持[gorm](https://gorm.io/gorm) 数据库组件、支持读写分离，数据库主从
7. 支持web界面 使用[ Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)
8. 支持支持多角色的RBAC权限控制，使用[casbin](https://github.com/casbin/casbin/v2)
9. 后续支持工具生成项目
10. 支持热编译[fresh](https://github.com/gravityblast/fresh)

### 联系作者
![my](https://activity-urt.oss-cn-beijing.aliyuncs.com/ecitic/%E6%88%91%E7%9A%84.jpg)
### 项目图片

![login](https://activity-urt.oss-cn-beijing.aliyuncs.com/ecitic/login.png)

![login](https://activity-urt.oss-cn-beijing.aliyuncs.com/ecitic/%E9%A6%96%E9%A1%B5.png)

![login](https://activity-urt.oss-cn-beijing.aliyuncs.com/ecitic/%E6%9D%83%E9%99%90%E9%85%8D%E7%BD%AE.png)
