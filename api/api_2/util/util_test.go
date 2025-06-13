package util

import (
	"fmt"
	"testing"
)

func Test_JsonConvert(t *testing.T) {
	type Sample[T int | string] struct {
		Name  string `json:"name"`
		Value T      `json:"value"`
	}

	va := Sample[int]{
		Name:  "sample",
		Value: 1,
	}

	formt_byte := StructToJson(va)
	fmt.Println(formt_byte, va)
	fmt.Println(ByteToStruct[Sample[int]](formt_byte))

}

func Test_env(t *testing.T) {
	SetEnv("sample", "sample_value")

	fmt.Println(GetEnv("sample", "not fund env name"))
}
