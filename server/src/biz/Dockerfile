# 使用官方的 golang 镜像作为构建环境
FROM golang:1.22.8 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码到工作目录
COPY . .

# 编译 Go 应用
RUN go build -o biz

# 使用更小的基础镜像
FROM my-custom-alpine

# 设置工作目录
WORKDIR /root/

# 从构建镜像中复制编译好的应用到运行环境中
COPY --from=builder /app/biz .

# 暴露应用的端口
EXPOSE 8888

# 运行应用, 会被docker-compose中的覆盖
CMD ["./biz"]
