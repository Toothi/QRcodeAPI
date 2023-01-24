package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"qrapi/tools"
	"time"
)

func Response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"QRcode API"}`))
	}
}

func Request_URL(a http.ResponseWriter, b *http.Request) {
	a.Header().Set("Content-Type", "application/json")
	a.WriteHeader(http.StatusOK)
	dataGet := b.URL.Query()
	Geturl := dataGet.Get("url")
	ip := b.RemoteAddr
	log.Printf("[%s] [请求生成]: %s", ip, Geturl)
	tools.Generate_QRcode(Geturl)
	filename := tools.Checks(Geturl)
	if filename == Geturl {
		fmt.Fprintf(a, `{"status":%d,GetUrl":"%s","DownloadPath":"/images/%s.png","QRcode_type":"text"}`, http.StatusOK, Geturl, filename)
	} else {
		fmt.Fprintf(a, `{"status":%d,GetUrl":"%s","DownloadPath":"/images/%s.png","QRcode_type":"url"}`, http.StatusOK, Geturl, filename)
	}
	time.AfterFunc(25*time.Second, func() { //异步执行匿名函数
		os.Remove("./images/" + filename + ".png")
		log.Println("[INFO]: 已清理图像")
	})
}
