# moss
moss is a simple and lightweight web content management system

------


## Get started
+ [Download program file](https://github.com/deep-project/moss/releases)
+ Run
> ./moss
+ Start successfully
> Use sqlite by default<br>
> Default management path /admin

------

### Configuration file(conf.toml)

| key  | Description       | Default   |
|------|-------------------|-----------|
| addr | listening address | random    |
| db   | database type     | sqlite    |
| dsn  | data source name  | ./moss.db |

+ Data source name Examples

| Type       | dsn Example                                                                        |
|------------|------------------------------------------------------------------------------------|
| sqlite     | ./data.db?_pragma=journal_mode(WAL)                                                |
| mysql      | user:password@tcp(127.0.0.1:3306)/moss?charset=utf8mb4&parseTime=True              |
| postgresql | host=127.0.0.1 port=5432 user=postgres password=123456 dbname=moss sslmode=disable |



### command
| key         | Description       | Example                                  |
|-------------|----------|----------------------------------------|
| --username  | reset administrator username |                                        |
| --password  | reset administrator password  |                                        |
| --adminpath | reset administration path    | ./moss --adminpath="admin"             |
| --config    | Define profile path | ./moss --config="/home/othername.toml" |

> ###### show more information by ./moss --help 