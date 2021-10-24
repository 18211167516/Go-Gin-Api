## 1、项目说明

> [!NOTE]
> 该项目为开源项目[Go-Gin-Api]项目文档

- 用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架
- 数据校验 `go-playground` gin框架内置，已支持错误信息转中文
- 数据库：采用`MySql`(5.6.44)版本，使用`gorm`实现对数据库的基本操作。
- API文档：使用`Swagger`构建自动化文档。
- 热重启`HotStart`,我自己写的一个
- 配置文件 使用`viper`解析
- 日志：使用`zap`实现日志记录。
- RBAC权限：使用`casbin`实现

