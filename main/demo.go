package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)


func init() {
	// 获取日志文件句柄
	// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`/data/go.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}


func index(w http.ResponseWriter, r *http.Request) {
	log.Println("我访问了Hello World这个路径")
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	log.Println("我访问了Health check这个路径")
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func check1(w http.ResponseWriter, r *http.Request) {
	log.Println("我访问了/api/v1这个路径")
	fmt.Fprintf(w, "<h1>我访问了/api/v1</h1>")
}

func main() {

	server := http.NewServeMux() // create a new mux server
	server.Handle("/metrics", promhttp.Handler()) // register a new handler for the /metrics endpoint
	server.HandleFunc("/", index)
	server.HandleFunc("/health_check", check)
	server.HandleFunc("/api/v1", check1)



	fmt.Println("Server starting...")
	//http.ListenAndServe(":3000", server)
	log.Fatal(http.ListenAndServe(":3000", server))



}
