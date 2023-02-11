# moss
moss is a simple and lightweight web content management system

moss是一个简单轻量的内容管理系统

------

+ [English document](https://github.com/deep-project/moss/blob/main/doc/README_EN.md)

## 开始使用
+ [下载程序文件](https://github.com/deep-project/moss/releases)
+ 运行
> ./moss
+ 启动成功
> 默认启动后使用sqlite<br>
> 默认后台地址 /admin

------

### 配置文件(conf.toml)

| key    | 说明     | 默认          |
|--------|--------|-------------|
| addr   | 监听地址   | 随机          |
| db     | 数据库类型  | sqlite      |
| dsn    | 数据源    | ./moss.db   |

+ 数据源示例

| Type       | dsn 示例                                                                             |
|------------|------------------------------------------------------------------------------------|
| sqlite     | ./data.db?_pragma=journal_mode(WAL)                                                |
| mysql      | user:password@tcp(127.0.0.1:3306)/moss?charset=utf8mb4&parseTime=True              |
| postgresql | host=127.0.0.1 port=5432 user=postgres password=123456 dbname=moss sslmode=disable |



### 命令行
| key         | 说明       | 示例                                     |
|-------------|----------|----------------------------------------|
| --username  | 重置管理员用户名 |                                        |
| --password  | 重置管理员密码  |                                        |
| --adminpath | 重置后台路径   | ./moss --adminpath="admin"             |
| --config    | 指定配置文件路径 | ./moss --config="/home/othername.toml" |

> ###### 可以通过 ./moss --help 查看更多信息