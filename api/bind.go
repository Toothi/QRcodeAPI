package api

import (
	"fmt"
	"log"
	"net/http"
	"qrapi/tools"
)

func Response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"QRcode API"}`))
	}
}

func Request_URL(a http.ResponseWriter, b *http.Request) {
	dataGet := b.URL.Query()
	Geturl := dataGet.Get("url")
	/* 调试输出
	fmt.Println("dataGet: ", dataGet)
	fmt.Println("Geturl: ", Geturl)
	*/
	log.Printf("[请求生成]: %s", Geturl)
	tools.Generate_QRcode(Geturl)
	fmt.Fprintf(a, `{"status":%d,GetUrl":"%s","DownloadUrl":"http://localhost:8080/images/qrcode.png"}`, http.StatusOK, Geturl)
}
