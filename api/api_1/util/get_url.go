package util

import (
	f "fmt"
)

func UrlCreate(base_url string, path_segments ...string) (result string) {
	result = base_url
	for _, path := range path_segments {
		result = f.Sprintf("%s/%s", result, path)
	}

	return result
}
