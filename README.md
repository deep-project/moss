# Moss
moss is a simple and lightweight web content management system

moss是一个简单轻量的内容管理系统

可以使用mysql、postgresql、sqlite数据库。后台支持12种语言，可切换明暗风格

使用中不懂的可以加群问我

QQ交流群：68396947

TG交流群：[https://t.me/mosscms](https://t.me/mosscms)


------

+ [English document](https://github.com/deep-project/moss/blob/main/docs/README_EN.md)
+ [主题制作](https://github.com/deep-project/moss/blob/main/docs/theme/README.md)
+ [模板文档](https://github.com/deep-project/moss/blob/main/docs/template/README.md)

+ [使用宝塔进程守护管理器部署程序](https://github.com/deep-project/moss/blob/main/docs/other/宝塔进程守护管理器部署程序.md)
+ [docker和docker-compose部署程序](./docs/other/docker和docker-compose部署程序.md)

![中文](https://user-images.githubusercontent.com/24670171/218475482-75030079-c2e3-4eb9-9f17-1713b15ad360.jpg)
![english](https://user-images.githubusercontent.com/24670171/218475496-4b2523b2-6bb6-43ac-a620-24f5ea0a5e3e.jpg)
### 暗色
![dark](https://user-images.githubusercontent.com/24670171/218475501-45527af5-c163-4331-b084-0c3943d6ff9c.jpg)
![list](https://user-images.githubusercontent.com/24670171/218475504-1ea5eb45-90cf-4810-aaa0-ca910b0165d5.jpg)


## 开始使用
+ [下载程序文件](https://github.com/deep-project/moss/releases)
+ 运行

      ./moss

+ 启动成功
> ##### 默认启动后使用sqlite<br>
> ##### 默认后台地址 /admin

------

### 配置文件(conf.toml)

| key  | 说明       | 默认      |
| ---- | ---------- | --------- |
| addr | 监听地址   | 随机      |
| db   | 数据库类型 | sqlite    |
| dsn  | 数据源     | ./moss.db?_pragma=journal_mode(WAL) |
      默认sqlite使用WAL方式打开，防止读取阻塞
+ 数据源示例

| Type       | dsn 示例                                                                           |
| ---------- | ---------------------------------------------------------------------------------- |
| sqlite     | ./data.db                                                |
| mysql      | user:password@tcp(127.0.0.1:3306)/moss?charset=utf8mb4&parseTime=True              |
| postgresql | host=127.0.0.1 port=5432 user=postgres password=123456 dbname=moss sslmode=disable |



### 命令行
| key         | 说明             | 示例                                   |
| ----------- | ---------------- | -------------------------------------- |
| --username  | 重置管理员用户名 |                                        |
| --password  | 重置管理员密码   |                                        |
| --adminpath | 重置后台路径     | ./moss --adminpath="admin"             |
| --config    | 指定配置文件路径 | ./moss --config="/home/othername.toml" |

> ###### 可以通过 ./moss --help 查看更多信息
