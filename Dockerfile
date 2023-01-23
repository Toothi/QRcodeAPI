FROM golang:alpine

ENV workpath /opt/QRcode-API
WORKDIR $workpath
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
EXPOSE 8080
ENTRYPOINT ["go","run","main.go"]
