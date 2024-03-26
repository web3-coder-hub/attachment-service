FROM golang:alpine AS builder

MAINTAINER "dev@devpeople.com"
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn" \
    TZ=Asia/Shanghai

ARG name
ARG env

WORKDIR $GOROOT/src/$name
COPY . .

ENV FLAGS "-w -s -X main.ENV=${env}"

COPY go.mod .
COPY go.sum .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories &&\
    apk update &&\
    apk add tzdata &&\
    cp /usr/share/zoneinfo/${TZ} /etc/localtime &&\
    go mod download &&\
    go build -ldflags "${FLAGS}" -o ${name} .

FROM ghcr.io/surnet/alpine-wkhtmltopdf:3.15.0-0.12.6-full

WORKDIR /gobuild
ARG name
ARG port
ARG swaggerDir
ENV NAME $name
ENV PORT $port
ENV SwaggerDir $swaggerDir

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories &&\
    apk update && apk add --no-cache \
    libstdc++ \
    libx11 \
    libxrender \
    libxext \
    libssl1.1 \
    ca-certificates \
    fontconfig \
    freetype \
    ttf-dejavu \
    ttf-droid \
    ttf-freefont \
    ttf-liberation \
    wget &&\
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub &&\
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-2.33-r0.apk &&\
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-bin-2.33-r0.apk  &&\
    wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.33-r0/glibc-i18n-2.33-r0.apk &&\
    apk add glibc-bin-2.33-r0.apk glibc-i18n-2.33-r0.apk glibc-2.33-r0.apk &&\
    apk add --no-cache --virtual .build-deps msttcorefonts-installer && update-ms-fonts &&\
    wget -q -O /usr/share/fonts/ttf-dejavu/simsun.ttf  https://github.com/bookc-man/alpine-wkhtmltopdf/raw/master/simsun.ttf &&\
    wget -q -O /usr/share/fonts/ttf-dejavu/simsun.ttc  https://github.com/bookc-man/alpine-wkhtmltopdf/raw/master/simsun.ttc &&\
    rm -rf *.apk && fc-cache -f && rm -rf /tmp/* && rm -rf /var/cache/apk/* &&\
    apk del .build-deps &&\
    /usr/glibc-compat/bin/localedef -i zh_CN -f UTF-8 zh_CN.UTF-8 &&\
    export LANG=zh_CN.UTF-8


COPY --from=builder $GOROOT/src/$NAME/$NAME /gobuild/
COPY --from=builder $GOROOT/src/$NAME/etc /gobuild/etc
COPY --from=builder $GOROOT/src/$NAME/internal/certs /gobuild/internal/certs
COPY --from=builder $GOROOT/src/$NAME/$SwaggerDir /gobuild/$SwaggerDir
COPY --from=builder /etc/localtime /etc/localtime

EXPOSE ${PORT}
ENTRYPOINT "./${NAME}" "-port=${PORT}" "-swagger-dir=${SwaggerDir}"