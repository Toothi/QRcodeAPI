FROM golang:alpine

ENV workpath /opt/QRcode-API
WORKDIR $workpath
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
EXPOSE 8080
ENTRYPOINT [$workpath"/start.sh"]
#ENTRYPOINT ["go","run","/opt/QRcode-API/main.go",">","log.txt","2>1&","&"]
