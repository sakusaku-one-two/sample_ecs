package service

import (
	"fmt"
	"testing"
)

const result_value = 200

func Test_sercie(t *testing.T) {
	i := 100
	if result_value == ReturnNumber(i) {
		fmt.Println("ok")
	}
}
