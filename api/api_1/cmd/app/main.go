package main

import (
	"api/env"
	"api/internal/service"
	"api/util"
	"fmt"
	"io"
	"net/http"
	"time"
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
		fmt.Println("/api/add ::: redis add data")
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		response, err := http.Post(util.UrlCreate(env.TARGET_SERVICE_URL, "add"), "application/json", r.Body)
		if err != nil {
			fmt.Println("add faild::", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Println("add faild::", err.Error())
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
			fmt.Println("all faild::", err.Error())
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

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/api/redis", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("/api/redis ::: 検証開始")

		res, err := http.Get(env.TARGET_SERVICE_URL)
		if err != nil {
			fmt.Println("そもそもredisserverと接続ができてない")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		real_ip, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("redis proxy serverがルートで返した値が取得できない")
		}

		fmt.Println(real_ip, "redis proxy serverのIP")

		redis_server_url := util.UrlCreate(env.TARGET_SERVICE_URL, "health")
		fmt.Println(redis_server_url, "ヘルスチェック実行")
		data, err := http.Get(redis_server_url)
		if err != nil {
			fmt.Println("redisとの接続に問題があり", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(data.StatusCode)
	})

	go func() {
		ticker := time.NewTicker(30 * time.Second)

		for range ticker.C {
			response, err := http.Get(env.TARGET_SERVICE_URL)
			if err != nil {
				fmt.Println("err =>", err.Error())
				continue
			}
			body_data, err := io.ReadAll(response.Body)
			response.Body.Close()
			if err != nil {
				fmt.Println("byte read ::", err.Error())
				continue
			}
			fmt.Println(string(body_data))

		}

		defer func() {
			ticker.Stop()
		}()

	}()

	err := http.ListenAndServe(env.SELF_SERVER_PORT, nil)
	fmt.Println("doen server", err)
}
