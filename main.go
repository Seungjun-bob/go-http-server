package main

import (
	"encoding/json"
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

	// get 요청 처리
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	// post 요청 처리
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// POST 메소드인지 확인
		if r.Method == http.MethodPost {
			// 요청 본문에서 JSON 데이터를 읽어오기
			var data struct {
				Name string `json:"name"`
			}

			// JSON 디코딩
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(w, "Invalid JSON data", http.StatusBadRequest)
				return
			}

			// 성공적으로 처리된 경우 응답
			fmt.Fprintf(w, "Received name: %s", data.Name)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// 2. 서버 시작
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
