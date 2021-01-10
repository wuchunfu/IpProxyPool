FROM golang:1.15-alpine as builder

# 设置环境变量, 指定编码
ENV LANG="en_US.UTF-8"

## 执行的命令
#RUN set -eux \
#    && go env -w GO111MODULE=on \
#    && go env -w GOPROXY=https://goproxy.cn,direct \
#    && go get github.com/wuchunfu/IpProxyPool

# 指定工作目录
WORKDIR /app

COPY . .

ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | \
    tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx \
    && chmod a+x /bin/upx

RUN set -eux \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o IpProxyPool . \
    && upx IpProxyPool

FROM alpine:3.12

# 指定镜像创建者信息
LABEL MAINTAINER="319355703@qq.com"

# 指定时区
ENV TIMEZONE Asia/Shanghai
ENV WORKPATH /app

# 指定工作目录
WORKDIR ${WORKPATH}

# 执行的命令
RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk upgrade \
    && apk update \
    && apk add --no-cache ca-certificates upx --no-progress bash tzdata busybox-extras \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo ${TIMEZONE} > /etc/timezone \
    && rm -rf /var/cache/apk/*

#COPY --from=builder /go/bin/IpProxyPool ${WORKPATH}
COPY --from=builder /app/IpProxyPool ${WORKPATH}
COPY ./dockerconfig/proxypool/conf/config.yaml ${WORKPATH}/conf/config.yaml

# 映射一个端口
EXPOSE 3000

ENTRYPOINT ["/app/IpProxyPool", "proxy-pool"]
