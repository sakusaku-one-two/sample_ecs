package main

import (
	"api/internal/service"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("POST")
			var body_byte []byte
			n, err := r.Body.Read(body_byte)
			if err != nil {
				fmt.Println("ERROR")
				return
			}

			fmt.Println(n, string(body_byte))
			w.WriteHeader(http.StatusOK)
			is_write_code, err := w.Write([]byte("hello!!!"))
			if err != nil {
				fmt.Println(is_write_code, err)
				return
			}
			return
		}

		if r.Method == "GET" {

			fmt.Println("GET")

			w.Write([]byte(service.SampleFunction()))

			return
		}

	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("doen server", err)
}
