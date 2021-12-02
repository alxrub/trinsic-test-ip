package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", test)
	router.NotFoundHandler = http.HandlerFunc(test2)
	http.ListenAndServe(":8080", router)
}

func test(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world"))
	fmt.Println("remote address:")
	fmt.Println(req.RemoteAddr)
	fmt.Println("forwarded for:")
	fmt.Println(req.Header.Get("X-Forwarded-For"))
}

func test2(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world not found"))
	fmt.Println("remote address:")
	fmt.Println(req.RemoteAddr)
	fmt.Println("forwarded for:")
	fmt.Println(req.Header.Get("X-Forwarded-For"))
}