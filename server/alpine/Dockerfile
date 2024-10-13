# 定义基础镜像为 Alpine
FROM golang:1.22.8-alpine3.20

# 安装gcompat
# gcompat is the go-to compatibility layer for Alpine users.
RUN apk add --no-cache gcompat

# 定义其他你想自定义的步骤
# ...

# 设置镜像的默认行为
CMD ["/bin/sh"]
