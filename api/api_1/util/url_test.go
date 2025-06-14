package util

import (
	"fmt"
	"testing"
)

func Test_url(t *testing.T) {
	fmt.Println(
		UrlCreate("base_url", "api", "user"),
	)
}
