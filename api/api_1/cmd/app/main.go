package main

import (
	"api/internal/service"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("POST")
			service.PostRetrun(w, r)
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
