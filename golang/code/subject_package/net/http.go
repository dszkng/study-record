package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//映射静态路由
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)

	//监听端口
	log.Fatal(http.ListenAndServe("0.0.0.0:1234", nil))
}

//解析 HTTP 报文
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	fmt.Fprintf(w, "%q", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

    for k,v := range r.Header {
	    fmt.Fprintf(w, "%q => %q\n", k, v)
    }
}
