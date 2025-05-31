package service

import (
	"strconv"
)

func ReturnNumber(arg_int int) int {
	return arg_int + arg_int
}

func ToString(target int) string {
	return strconv.Itoa(target)
}

func SampleFunction() string {
	return ToString(ReturnNumber(100))
}
