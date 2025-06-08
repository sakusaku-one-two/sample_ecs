package service

import (
	"encoding/json"
	"net/http"
)

type ResponseType struct {
	Name string `json:"name"`
}

func NewResonseType(url_address string) *ResponseType {
	return &ResponseType{
		Name: url_address,
	}
}

func ToByte[T any](arg *T) ([]byte, error) {
	return json.Marshal(*arg)
}

func ByteTo[T any](arg []byte) (T, error) {
	var data T

	err := json.Unmarshal(arg, &data)

	return data, err

}

func PostRetrun(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI
	w.Header().Set("Content-Type", "application/json")

	json_byte, err := ToByte[ResponseType](NewResonseType(url))

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("失敗しました"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json_byte)

}
