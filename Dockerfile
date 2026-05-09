# 只用一个镜像：golang 官方 Alpine 版
FROM golang:1.23-alpine

# 安装依赖
RUN apk add --no-cache git

# ==============================
# 1. 编译目录：/app/build
# ==============================
WORKDIR /app/build

# 拷贝 go mod 缓存依赖
COPY main/go.mod main/go.sum ./
# ENV GOPROXY=https://goproxy.cn,direct
# ENV GOSUMDB=off

# 下载依赖
RUN go mod download

# 拷贝源码
COPY main/ .

# 编译：输出到【运行目录 /app/run/app】
RUN mkdir -p /app/run \
    && CGO_ENABLED=0 go build -o /app/run/app ./cmd/web

# ==============================
# 2. 运行目录：/app/run
# ==============================
WORKDIR /app/run

# 暴露端口
EXPOSE 3000

# 运行【运行目录】里的二进制
CMD ["./app", "-a", ":3000"]