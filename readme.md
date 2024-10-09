
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

### docker部署

构建项目镜像
```bash
# 拉取源代码
git clone https://github.com/moreKing/web-firewall.git
cd web-firewall
# 构建镜像
docker build -f ./server/manifest/docker/Dockerfile  -t web-firewall:latest .
```

docker使用必须以特权和host模式运行才能操作主机的网络配置
```bash
docker run -itd --network host --privileged \
-v /etc/sysctl.conf:/etc/sysctl.conf \
-v  /proc:/host_proc \
web-firewall
```

数据持久化需要映射以下几个目录

```bash
-v /path/config:/web-firewall/manifest/config  # 配置文件 docker模式下 默认数据库文件也在此目录
-v /path/log:/web-firewall/log  # 日志文件目录 所有日志文件均会在此目录下
-v /path/resources/template:/web-firewall/resources/template  # 模板文件，用户可以自定义邮件样式
```

`/path/config`目录内需要提供config.yaml配置文件和db.sqlite3初始化数据库文件，可先运行一个镜像从里面对应位置拷贝出里面的默认文件，然后修改配置文件即可

### 项目截图

![登录](./img/login.png)

![主题](./img/theme.png)

![国际化](./img/i18.png)

![添加策略](./img/add_policy.png)

![添加流控](./img/add_limit.png)

![暗模式](./img/dark.png)

![自适应](./img/mobile.png)



## 捐赠支持

为了开源项目更好的发展，现在接受捐赠。如果您觉得本项目对您有所帮助，请通过下列方式进行捐赠。您的帮助将给我们动力持续更新，谢谢！

![收款码](./img/pay.png)