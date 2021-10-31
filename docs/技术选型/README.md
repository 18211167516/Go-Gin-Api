## 1、项目说明

> [!NOTE]
> 该项目为开源项目[Go-Gin-Api]项目文档

- 用`Gin`快速搭建基础restful风格API，`Gin`[文档](https://gin-gonic.com/zh-cn/docs/)是一个go语言编写的Web框架
- 数据校验 `go-playground` [github](https://github.com/go-playground/validator) gin框架内置，已支持错误信息转中文
- 数据库：采用`MySql`(5.6.44)版本，使用`gorm` [文档](https://gorm.io/zh_CN/)实现对数据库的基本操作。
- API文档：使用`Swagger`[github](https://github.com/swaggo/gin-swagger)构建自动化文档。
- 热重启`HotStart` [github](https://github.com/18211167516/hotstart),我自己写的一个
- 配置文件 使用`viper` [github](https://github.com/spf13/viper)解析
- 日志：使用`zap` [github](https://go.uber.org/zap)实现日志记录。
- RBAC权限：使用`casbin`[文档](https://casbin.org/docs/zh-CN/)实现
- 安全Cookie:使用`encrypt`[github](https://github.com/18211167516/encrypt)实现加密解密

