package env

import (
	"strings"
)

func Bool(key string, canSkip ...bool) bool {
	v := getEnv(key, canSkip...)
	return v == "true" || v == "1"
}

func BoolList(key string, canSkip ...bool) []bool {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]bool, len(vs))
	for index, v := range vs {
		result[index] = v == "true" || v == "1"
	}
	return result
}
