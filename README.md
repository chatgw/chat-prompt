# Chat-prompt

chat-prompt 是一个平台，旨在分享 DALL·E、Midjourney、Stable Diffusion、GPT 和 ChatGPT 提示词的分享平台. 它将是一个专为创造各种提示的 Prompt Engineer 和开发人员提供的提示词。本网站讲列出了最受欢迎的提示，用户浏览共享他们的提示词. 此外，我们还鼓励大家相互交流，互相改进大家的提示工程技能.


## 平台简介

* 前端技术栈 TypeScript、Vue3、Element-Plus、Vite、Pinia
* 后端技术栈 Golang、Gin、Gorm、MySQL、Redis、Casbin

## 内置功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。

## 运行

默认账号密码

admin/123456

克隆项目
`git clone https://github.com/airdb/chat-prompt.git`

### 前端

```bash
# 配置
1. 一键安装 .vscode 目录中推荐的插件
2. node 版本 17+
3. pnpm 版本 7.x

# 进入项目目录
cd web

# 安装依赖
pnpm i

# 启动服务
pnpm dev

# 预览预发布环境
pnpm preview:stage

# 预览正式环境
pnpm preview:prod

# 构建预发布环境
pnpm build:stage

# 构建正式环境
pnpm build:prod

# 代码格式化
pnpm lint
```

### 后端

```bash
# 配置
1. Go >= 1.19

# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

## 一键安装

Docker Compose 版本需要 V2

```bash
git clone https://github.com/airdb/chat-prompt
cd chat-prompt
docker compose -f docker-compose/docker-compose.yml build
docker compose -f docker-compose/docker-compose.yml up -d
```

浏览器打开 `http://ip:8500`
