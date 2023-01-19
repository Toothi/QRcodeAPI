package main

import (
	"log"
	"net/http"
	"qrapi/api"
)

func main() {
	http.HandleFunc("/", api.Response)
	http.HandleFunc("/qrcode", api.Request_URL)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))
	log.Println("Running at [:8080]")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
