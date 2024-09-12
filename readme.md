
# 简介

`web-firewall`基于golang+vue3 开发的Web Linux防火墙，前端使用`SoybeanAdmin`框架，后端使用`goframe2`，数据库支持 `sqlite3(默认)`/`postgresql` ，它可以在Linux系统中基于`nfatables`用于替代`firewalld`工具。

Gitee Star：[![gitee star](https://gitee.com/moujun/web-firewall/badge/star.svg)](https://gitee.com/moujun/web-firewall)

Github Star：[![github star](https://img.shields.io/github/stars/moreKing/web-firewall)](https://github.com/moreKing/web-firewall)

Github Forks：[![github forks](https://img.shields.io/github/forks/moreKing/web-firewall)](https://github.com/moreKing/web-firewall)

该防火墙可以提供以下功能



### 功能设计

   - [x] 本地策略
     - [x] 出站策略 output链
     - [x] 出站流控 output链
     - [x] 入站策略 input链
     - [x] 入站流控 input链
   - [x] 路由策略
     - [x] DNAT prerouting链
     - [x] SNAT postrouting
     - [x] 转发策略forward 链
     - [x]  流量控制 forward 链
   - [x] 审计
     - [x] 登录日志
     - [x] 配置日志
     - [x] webSSH日志
   - [x] WebShell
     - [x] webshell 支持rzsz
     - [ ] 文件上传下载
     - [ ] 文本在线编辑
   - [ ] 首页
   - [x] 系统设置
     - [ ] ip黑名单 prerouting链
     - [x] 本地密码
     - [x] 会话配置
     - [x] 邮件配置
     - [x] 短信配置
     - [x] 登录设置

### 安装

本项目提供一个已经打包编译好的项目，用户仅需自己[下载](https://moujun.top/web-firewall/version.html)本项目解压后，执行 里面的`install.sh`文件即可，如果自己编译项目请根据前后端代码自行进行打包即可

```shell
unzip v1.1.0.zip
cd v1.1.0
bash install.sh

# 查看服务是否正常运行
systemctl status web-firewalld

# 建议停用firewalld服务
systemctl disable  firewalld
systemctl stop  firewalld
```

访问地址：http://ip:8000

默认账号密码：admin/admin


### 项目截图

![登录](./img/login.png)

![主题](./img/theme.png)

![国际化](./img/i18.png)

![添加策略](./img/add_policy.png)

![添加流控](./img/add_limit.png)

![暗模式](./img/dark.png)

![自适应](./img/mobile.png)





