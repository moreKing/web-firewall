
# 简介

本项目是一个基于golang+vue3 开发的Web防火墙，它可以在Linux系统中替代firewalld工具。该防火墙可以提供以下功能：

1. 查看防火墙状态
2. 开放端口
3. 关闭端口
4. 查看防火墙规则
5. 添加防火墙规则
6. 删除防火墙规则
7. 查看防火墙日志
8. 查看防火墙统计信息

##  安装

1. 克隆项目到本地

```bash
git clone https://github.com/yourname/firewall.git
```

2. 进入项目目录

```bash
cd firewall
```

3. 安装依赖

```bash
go mod tidy
```

4. 编译项目

```bash
go build
```

5. 启动项目

```bash
./firewall
```

6. 访问Web界面

在浏览器中访问`http://localhost:8080`即可访问Web界面。





## 功能设计

1. 菜单
   - 本地策略
     - [ ] 出站策略 output链
     - [ ] 入站策略 input链
   - [ ] 地址转换（NAT）
     - [ ] DNAT prerouting链
     - [ ] SNAT postrouting
   - [ ] 转发策略（作为网关时）forward 链
   - [ ] 连接数/流量 控制 input链/forward 链
   - [ ] ip黑名单 prerouting链
   - 审计
     - [ ] 登录日志
     - [ ] 配置日志
     - [ ] webSSH日志
   - WebSSH
     - [ ] ssh操作
     - [ ] 文件上传下载
     - [ ] 文本在线编辑
   - 首页
   - 系统设置
     - [x] 本地密码
     - [x] 会话配置
     - [x] 邮件配置
     - [x] 短信配置
     - [x] 登录设置

