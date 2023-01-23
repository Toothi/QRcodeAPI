package tools

import (
	"github.com/skip2/go-qrcode"
	"strings"
)

func Checks(a string) string {
	checkd1 := strings.Contains(a, "http://")
	checkd2 := strings.Contains(a, "https://")
	if checkd1 == true {
		rename := string(a[7:])
		return rename
	} else if checkd2 == true {
		rename := string(a[8:])
		return rename
	} else {
		return "ERROR"
	}
}

func Generate_QRcode(parameters string) {
	rename := Checks(parameters)
	qrcode.WriteFile(parameters, qrcode.Medium, 256, "./images/"+rename+".png")
}
