# ACG.Fate

acgfate.com 是一个基于 [gin](https://github.com/gin-gonic/gin) 实现的论坛类网站。

## 项目结构

```
.
├───api
│   └───http    // http API
│       └───v1  
├───cache       // 缓存
├───config      // 配置文件
├───dao         // 数据库操作层
├───docs        // Swagger
├───middleware  // 中间件
├───model       // 数据库模型
├───router      // 路由
├───serializer  // 用户响应序列化
├───service     // 服务
└───util        // 独立组件
```

## 第三方库

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://gorm.io/)
- [go-redis](https://github.com/go-redis/redis)
- [viper](https://github.com/spf13/viper)
- [validator](https://github.com/go-playground/validator)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [zap](https://github.com/uber-go/zap)

## TODO

- 优化 SQL 语句
- 参数校验
- 记录日志
- 实现审核机制
