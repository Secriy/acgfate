ACG.Fate 类 Reddit 论坛项目。

## 功能

- 用户登录、注册、个人信息查询
- 投稿分区查询、列表查询
- 文章投稿
- 文章查询、列表查询
- 分页查询
- 文章点赞
- 文章热度榜单

## 接口

```
POST   /api/v1/user/register     --> acgfate/api/http/v1.UserRegister 
POST   /api/v1/user/login        --> acgfate/api/http/v1.UserLogin 
GET    /api/v1/category/:name    --> acgfate/api/http/v1.CategoryDetail 
GET    /api/v1/category/list     --> acgfate/api/http/v1.CategoryList 
GET    /api/v1/word/list         --> acgfate/api/http/v1.WordList 
GET    /api/v1/word/:id          --> acgfate/api/http/v1.WordDetail 
GET    /api/v1/word/trend        --> acgfate/api/http/v1.WordTrend 
GET    /api/v1/user/info         --> acgfate/api/http/v1.UserInfo 
POST   /api/v1/word/post         --> acgfate/api/http/v1.WordPost 
DELETE /api/v1/word/:id/delete   --> acgfate/api/http/v1.WordDelete 
PUT    /api/v1/word/:id/like     --> acgfate/api/http/v1.WordLike 
DELETE /api/v1/word/:id/like     --> acgfate/api/http/v1.WordUnlike 
```

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

## 框架 & 第三方库

以下是本项目主要使用的第三方库：

- [gin](https://github.com/gin-gonic/gin)：轻量级 web 框架
- [sqlx]()：SQL 简单封装库
- [go-redis](https://github.com/go-redis/redis)：Redis 操作库
- [viper]()：配置文件加载
- [validator]()：参数校验
- [jwt-go]()：JWT 库
- [zap](https://go.uber.org/zap)：日志库
- [gin-swagger](https://github.com/swaggo/gin-swagger)：Swagger RESTFul 文档生成

