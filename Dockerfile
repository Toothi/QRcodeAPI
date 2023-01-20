FROM golang:alpine

WORKDIR /opt/QRcode-API
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
EXPOSE 8080
ENTRYPOINT ["go","run","/opt/QRcode-API/main.go"]
