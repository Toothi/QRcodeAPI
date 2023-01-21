# QRcodeAPI
Web application interface for generating QR code
## Usage:
```text
http://0.0.0.0:8080/qrcode?url=http://google.com
http://0.0.0.0:8080/images/qrcode.png
```
## Docker
```text
git clone https://github.com/Toothi/QRcodeAPI
cd QRcodeAPI
docker build -t qrcode .
docker run -d -P --name qrcode-api qrcode
```
## Server
### Linux
```shell
go mod tidy
go run main.go
```
### Windows
##### Download Releases
## Build
### Linux and Windows
```shell
go mod tidy
go build main.go
```
