package env

import (
	u "api/util"
	"fmt"
)

var (
	SELF_SERVER_PORT    string
	SERVICE_NAME        string
	NAME_SPACE          string
	TARGET_SERVICE_PORT string
	TARGET_SERVICE_URL  string
)

func init() {
	// set up
	SELF_SERVER_PORT = fmt.Sprintf(":%s", u.GetEnv("SELF_SERVER_PORT", "8081"))
	TARGET_SERVICE_PORT = u.GetEnv("TARGET_SERVICE_PORT", "8080")
	SERVICE_NAME = u.GetEnv("SERVICE_NAME", fmt.Sprintf("localhost:%s", TARGET_SERVICE_PORT))
	NAME_SPACE = u.GetEnv("NAME_SPACE", "")

	if NAME_SPACE != "" {
		TARGET_SERVICE_URL = fmt.Sprintf("http://%s.%s:%s", SERVICE_NAME, NAME_SPACE, TARGET_SERVICE_PORT)
	} else {
		TARGET_SERVICE_URL = fmt.Sprintf("http://localhost:%s", TARGET_SERVICE_PORT)
	}

}
