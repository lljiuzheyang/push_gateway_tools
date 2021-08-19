FROM golang:1.16 as builder
ENV GOPROXY="https://goproxy.io"
ENV GOPATH=$HOME/go
ENV GOBIN=$HOME/go/bin
ENV PATH=$PATH:$GOPATH/bin
COPY . /app/
# 下载指定的包，go.mod已经记录，可以直接使用
RUN cd /app && go build -o pushGateway ./main.go
RUN cd /app && go build -o shellGateway ./shell.go

FROM ubuntu:20.10

WORKDIR /
COPY --from=builder /app/pushGateway /
COPY --from=builder /app/shellGateway /
COPY --from=builder /app/conf /conf
COPY --from=builder /app/start.sh /
ENV GIN_MODE=release

CMD ["/shellGateway"]