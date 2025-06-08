package service

import (
	"fmt"
	"testing"
)

func Test_return(t *testing.T) {
	sample := NewResonseType("my_address")

	fmt.Println(sample)
	byte_data, err := ToByte[ResponseType](sample)
	fmt.Println(byte_data, err)
	re_data, err := ByteTo[ResponseType](byte_data)
	fmt.Println(re_data, err)
}
