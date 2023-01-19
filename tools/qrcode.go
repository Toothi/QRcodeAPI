package tools

import (
	"github.com/skip2/go-qrcode"
)

func Generate_QRcode(parameters string) {
	qrcode.WriteFile(parameters, qrcode.Medium, 256, "./images/qrcode.png")
}
