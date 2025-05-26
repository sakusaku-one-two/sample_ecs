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

	w.Write([]byte{"sdfadsfsasf"})

}
