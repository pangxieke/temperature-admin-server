FROM registry.cn-shenzhen.aliyuncs.com/yunzhimeng/portrays-builder:latest as builder
WORKDIR /go/src/temp-admin
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
ENV GO111MODULE on
RUN go mod init
COPY . .

RUN CGO_ENABLED=0 go build -o app_d ./main.go

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
LABEL \
    SERVICE_80_NAME=temp-admin_http \
    SERVICE_NAME=temp-admin \
    description="测温后台" \
    maintainer="***"

EXPOSE 19980
COPY --from=builder /go/src/temp-admin/app_d /bin/app
ENTRYPOINT ["app"]
