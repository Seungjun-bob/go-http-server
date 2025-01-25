package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. 기본 핸들러 함수 정의
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! This is a simple HTTP server in Go.")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is the /hello endpoint!")
	})

	// 2. 서버 시작
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
