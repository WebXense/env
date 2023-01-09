package env

import (
	"strings"
)

func String(key string, canSkip ...bool) string {
	return getEnv(key, canSkip...)
}

func StringList(key string, canSkip ...bool) []string {
	return strings.Split(getEnv(key, canSkip...), ",")
}
