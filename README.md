# ICU Blog System

ICU Blog 是一个现代化的博客系统，集成了 AI 聊天功能的 Web 应用程序。它提供了完整的用户管理、文章发布、评论互动等功能，并创新性地集成了 AI 聊天功能。

## 主要功能

### 用户系统

- 用户注册和登录
- 个人信息管理（头像、签名等）
- JWT 认证
- 验证码支持

### 文章系统

- 文章发布和管理
- 标签系统
- 文章图片上传
- 富文本编辑支持

### 评论系统

- 文章评论
- 评论管理
- 用户互动

### AI 聊天

- 基于 OpenAI API 的 AI 聊天功能
- 实时聊天（SSE）支持
- 智能对话

### 其他功能

- 文件上传和管理
- 实时消息通知
- 标签管理

## 技术栈

### 后端

- Go 语言
- Gin Web 框架
- GORM 数据库 ORM
- JWT 认证
- OpenAI API 集成
- Server-Sent Events (SSE)

### 数据库

- MySQL

### 工具和库

- Viper (配置管理)
- Logrus (日志管理)
- Captcha (验证码)

## 项目结构

```
.
├── cmd/
│   └── app/            # 应用程序入口
├── config/             # 配置文件
├── internal/
│   ├── controller/     # HTTP 请求处理
│   ├── model/         # 数据模型
│   ├── repository/    # 数据访问层
│   ├── route/         # 路由配置
│   ├── service/       # 业务逻辑
│   └── utils/         # 工具函数
└── scripts/           # 脚本文件
```

## 快速开始

### 前置要求

- Go 1.16+
- MySQL 5.7+
- OpenAI API 密钥（用于 AI 聊天功能）

### 配置

1. 复制配置文件模板并修改配置
2. 设置数据库连接信息
3. 配置 OpenAI API 密钥

### 运行

```bash
# 构建
go build -o app cmd/app/main.go

# 运行
./app
```

## API 文档

### 用户相关

- POST /api/user/register - 用户注册
- POST /api/user/login - 用户登录
- GET /api/user/info - 获取用户信息
- PUT /api/user/info - 更新用户信息

### 文章相关

- POST /api/article - 发布文章
- GET /api/article - 获取文章列表
- GET /api/article/:id - 获取文章详情
- PUT /api/article/:id - 更新文章
- DELETE /api/article/:id - 删除文章

### 评论相关

- POST /api/comment - 发表评论
- GET /api/comment - 获取评论列表
- DELETE /api/comment/:id - 删除评论

### AI 聊天

- GET /api/chat/stream - 建立 AI 聊天连接

## 许可证

[MIT License](LICENSE)
