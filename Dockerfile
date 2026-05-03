FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

# 拷贝整个项目
COPY main/ .

# 下载依赖
RUN go mod download

# 编译
RUN go build -o app ./cmd/web

# 暴露端口
EXPOSE 3000

# 运行
CMD ["./app", "-addr", ":3000"]