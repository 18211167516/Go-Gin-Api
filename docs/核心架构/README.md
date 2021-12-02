
> [!TIP]
> 整体架构


## 1、目录结构
```
    ├─app  	     （项目核心目录）
    |  ├─controller     （控制器）
    |  ├─middleware     （中间件）
    |  ├─models         （数据结构层）
    |  ├─request        （数据请求层，定义特殊请求结构体以及数据校验）
    |  ├─request        （数据展示层定义结构体）
    |  ├─services       （服务层）
    ├─config            （配置包）
    ├─core  	        （內核、项目核心库）
    ├─docs  	        （swagger文档目录）
    ├─global            （全局变量）
    ├─initialize         (初始化)  
    ├─routes            （路由）
    ├─static            （静态文件包括config目录）
    ├─templates         （模板目录）
    ├─tests             （测试）
    └─tool	            （工具包）

```

## 2、组件支持

1. 支持[Swagger](https://github.com/swaggo/gin-swagger)接口文档生成
2. 支持[jwt](https://github.com/golang-jwt/jwt)鉴权
3. 支持[logrus](https://github.com/sirupsen/logrus)(准备替换成zap)
4. 支持[viper](https://github.com/spf13/viper) 配置文件解析
5. 支持go1.6.0 go:embed特性,打包包含静态文件
6. 支持[gorm](https://gorm.io/gorm) 数据库组件、支持读写分离，数据库主从
7. 支持web界面 使用[ Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)
8. 支持支持多角色的RBAC权限控制[支持按钮级权限控制]，使用[casbin](https://github.com/casbin/casbin/v2)
9. 支持工具快速生成`controller`、`model`、`service`、`view` 
10. 支持热编译[fresh](https://github.com/gravityblast/fresh)