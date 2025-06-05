package service

import (
	"encoding/json"
	"net/http"
)

type ResponseType struct {
	name string
}

func NewResonseType(url_address string) *ResponseType {
	return &ResponseType{
		name: url_address,
	}
}

func ToByte[T any](arg *T) ([]byte, error) {
	return json.Marshal(arg)
}

func PostRetrun(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI

	json_byte, err := ToByte[ResponseType](NewResonseType(url))

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("失敗しました"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json_byte)

	return

}
