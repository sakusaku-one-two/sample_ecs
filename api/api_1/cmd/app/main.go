package main

import (
	"api/env"
	"api/internal/service"
	"api/util"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

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

	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		// redisに追加を行う　POSTのみ
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		response, err := http.Post(util.UrlCreate(env.TARGET_SERVICE_URL, "add"), "application/json", r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)

	})

	http.HandleFunc("/api/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		fmt.Println(util.UrlCreate(env.TARGET_SERVICE_URL, "all"))

		resp, err := http.Get(util.UrlCreate(env.TARGET_SERVICE_URL, "all"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(data)

	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	err := http.ListenAndServe(env.SELF_SERVER_PORT, nil)
	fmt.Println("doen server", err)
}
