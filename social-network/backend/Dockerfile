# 使用Go官方镜像作为构建环境
FROM golang:1.22 as builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.* ./
# 下载依赖
RUN go mod download

# 复制项目文件到工作目录
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用scratch作为最终镜像的基础，创建一个最小化的容器
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建者镜像中复制构建好的应用到当前目录
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 8888

# 运行应用
CMD ["./main"]
