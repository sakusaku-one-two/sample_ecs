package handler

import (
	"fmt"
	"net/http"
)

func ReflectMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		fmt.Println("POST以外のメソッドが呼ばれました", r.Method)
		return
	}
	var body_byte []byte
	_, err := r.Body.Read(body_byte)
	if err != nil {
		w.Write([]byte("読み取りに失敗しました"))

		return
	}
	w.Write(body_byte)

}
