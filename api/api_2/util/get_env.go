package util

import (
	"os"
)

// ----------------enviroment--------------------//

func SetEnv(env_name, env_value string) {
	os.Setenv(env_name, env_value)
}

func GetEnv(env_name, defalut string) string {
	reuslt, ok := os.LookupEnv(env_name)
	if ok {
		return reuslt
	}
	return defalut
}
