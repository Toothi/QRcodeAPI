FROM golang:1.18-alpine

ENV workpath /opt/QRcode-API
WORKDIR $workpath
COPY go.mod .
COPY go.sum .
RUN go env -w GO111MODULE=on \
&&  go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
COPY . .
EXPOSE 8080
ENTRYPOINT ["go","run","main.go"]
