# 单镜像
FROM golang:1.23-alpine

RUN apk add --no-cache git

# ==========================
# 编译目录（临时用）
# ==========================
WORKDIR /app/build

COPY main/go.mod main/go.sum ./
RUN go mod download

COPY main/ .

# 编译：把可执行文件放在 /app/main（运行目录外面）
RUN CGO_ENABLED=0 go build -o /app/main ./cmd/web

# ==========================
# 运行目录（空的，用于映射）
# ==========================
WORKDIR /app/run
VOLUME /app/run

# 端口
EXPOSE 3000

CMD ["/app/main", "-a", ":3000"]