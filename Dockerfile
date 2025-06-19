FROM golang:alpine AS builder

# 设置环境变量
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# 替换 Alpine 源为阿里云镜像
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache git

WORKDIR /build
COPY . .

# 自动修复所有 Go 文件中的错误注释
RUN find . -name "*.go" -type f -exec sed -i '/direct #/d' {} \; && \
    find . -name "*.go" -type f -exec sed -i '/direct%20#/d' {} \;

RUN go mod download && \
    go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/main /app/

# 设置时区和安装必要依赖
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

CMD ["./main"]