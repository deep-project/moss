# 目录

1. [安装dokcer和docker-compose](#1-安装dokcer和docker-compose)
   1.1 [关于docker和docker-compose的安装](#11-关于docker和docker-compose的安装)
   1.2 [docker安装](#12-docker安装)
   1.3 [docker-compose安装](#13-docker-compose安装)
   1.4 [docker更换镜像加速器](#14-docker更换镜像加速器)

2. [创建运行镜像](#2-创建运行镜像)
   2.1 [创建原始目录](#21-创建原始目录)
   2.2 [创建Dockerfile](#22-创建Dockerfile)
   2.3 [创建所需文件](#23-创建所需文件)
   2.4 [创建镜像](#24-创建镜像)

3. [运行容器](#3-运行容器)
   3.1 [docker运行](#31-docker运行)
   3.2 [docker-compose运行](#32-docker-compose运行)

## 1. 安装dokcer和docker-compose
### 1.1 关于docker和docker-compose的安装
以`Debian`系统为例，其他系统对应更改包管理器命令。
- Debian/Ubuntu: apt-get
- RedHat/CentOS: yum

### 1.2 docker安装
- 大陆外
```
apt-get update -y && apt-get install -y curl wget && curl -fsSL https://get.docker.com | bash -s docker
```
- 国内
```
apt-get update -y && apt-get install -y curl wget && curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

### 1.3 docker-compose安装
- 大陆外
```
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose &&  chmod +x /usr/local/bin/docker-compose
```
- 国内
```
curl -L "https://ghproxy.com/https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose &&  chmod +x /usr/local/bin/docker-compose
```
如果 `/usr/local/bin` 不在环境变量 PATH 里
```
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```
检查`docker-compose`版本
```
docker-compose --version
```

### 1.4 docker更换镜像加速器

`hub.docker.com`似乎被DNS污染了，

大陆机子推荐更换镜像加速器。非大陆机子不推荐。

关于源的有效性，可以参考以下这个开源项目，进行更改

- https://gist.github.com/y0ngb1n/7e8f16af3242c7815e7ca2f0833d3ea6

终端运行
```
echo >/etc/docker/daemon.json
cat>/etc/docker/daemon.json <<END
{
  "registry-mirrors": [
    "https://hub-mirror.c.163.com",
    "https://docker.m.daocloud.io",
    "https://ghcr.io",
    "https://mirror.baidubce.com",
    "https://docker.nju.edu.cn"
  ]
}
END
systemctl restart docker
```


## 2. 创建运行镜像

### 2.1 创建原始目录
```
mkdir -p moss
```
### 2.2 创建Dockerfile

可执行程序连接按需修改，下载链接在[release页面](https://github.com/deep-project/moss/releases)获取

```
https://github.com/deep-project/moss/releases/download/v0.1.1/moss_linux_amd64
```
当国内因为网络问题无法连接`github`时，可以使用`ghproxy`等进行加速
```
https://ghproxy.com/https://github.com/deep-project/moss/releases/download/v0.1.1/moss_linux_amd64
```
- 创建ockerfile

以下是个完整命令复制到终端运行

```
cat > ./moss/Dockerfile <<EOF
# 使用 golang:alpine3.18 作为基础镜像，该镜像包含了 Golang 环境和 Alpine 3.18 发行版
FROM golang:alpine3.18

# 使用中国科学技术大学的镜像源来替换默认的 Alpine 软件包镜像源
RUN echo "https://mirrors.ustc.edu.cn/alpine/v3.18/main" > /etc/apk/repositories \
    && echo "https://mirrors.ustc.edu.cn/alpine/v3.18/community" >> /etc/apk/repositories

# 更新软件包索引，升级已安装的软件包，并安装 wget
RUN apk update && apk upgrade && \
    apk add --no-cache wget && \
    mkdir -p /app && \
    mkdir -p /app/mosscms && \
    wget -O /app/moss https://github.com/deep-project/moss/releases/download/v0.1.1/moss_linux_amd64 && \
    chmod +x /app/moss

# 安装时区数据包，并将时区设置为 Asia/Shanghai
RUN apk --no-cache add tzdata
RUN echo "Asia/Shanghai" > /etc/timezone

# 设置工作目录为 /app/mosscms
WORKDIR /app/mosscms

# 将容器的 9008 端口暴露出来，允许外部访问
EXPOSE 9008

# 定义容器启动时执行的默认命令为 /app/moss
CMD ["/app/moss"]

EOF
```
### 2.3 创建所需文件

- 创建配置文件目录
```
mkdir -p ./moss/data
```
- 创建配置文件

以下是个完整命令复制到终端运行
```
cat > ./moss/data/conf.toml <<EOF
addr = ':9008'
db = 'sqlite'
dsn = './moss.db?_pragma=journal_mode(WAL)'

EOF

```

### 2.4 创建镜像
- 进入目录
```
cd moss
```
- 创建镜像
```
# docker build: 用于构建 Docker 镜像的命令
# -t moss:0.1.1: 指定构建的镜像名称为 moss，标签为 0.1.1
# .: 构建上下文路径，Docker 将在该路径下查找 Dockerfile 和相关资源来构建镜像
docker build -t moss:0.1.1 .

```
## 3. 运行容器
### 3.1 docker运行

```
docker run -d \
--name moss \
-p 9008:9008 \
-v ./data/conf.toml:/app/mosscms/conf.toml \
-v ./data:/app/mosscms \
moss:0.1.1
```

通过`http://IP:9008`进行访问

### 3.2 docker-compose运行

- 创建docker-compose.yml

以下是个完整命令复制到终端运行
```
cat >docker-compose.yml <<EOF
version: '3'
services:
  moss:
    image: moss:0.1.1
    container_name: moss
    ports:
      - "9008:9008"
    volumes:
      - ./data/conf.toml:/app/mosscms/conf.toml
      - ./data:/app/mosscms
    restart: always
EOF
```
- 运行
```
docker-compose up -d
```
通过`http://IP:9008`进行访问