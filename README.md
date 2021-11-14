## 框架 & 库

- [gin](https://github.com/gin-gonic/gin)
- [gorm]()
- [go-redis](https://github.com/go-redis/redis)
- [viper]()
- [validator]()
- [jwt-go]()
- [zap]()

## 目录结构

```
.
├── api
│   └── http        // HTTP API
│       └── v1
├── cache           // 缓存
├── config          // 项目配置操作
├── database        // 数据库层
├── docs            // Swagger 文档
├── middleware      // web 中间件
├── model           // 数据模型
├── router          // 路由
├── schema          // 数据库结构 SQL 文件
├── serializer      // 响应序列化器
├── service         // 服务层 
└── util            // 独立工具
    ├── logger      // 日志组件
    └── snowflake   // 雪花算法组件
```
