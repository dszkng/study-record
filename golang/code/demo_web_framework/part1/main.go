package main

import (
	"example/gin"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := gin.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "%q => %q\n", k, v)
		}
	})

	log.Fatal(r.Run(":1234"))
}
