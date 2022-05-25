FROM golang:alpine

ENV GO111MODULE=on\
    CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64\
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/cqupt-post/keycenter

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /home/cqupt-post/keycenter .

RUN mkdir src .

EXPOSE 8901

CMD ["/dist/app"]