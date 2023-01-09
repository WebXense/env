package env

import (
	"strconv"
	"strings"
)

func Int(key string, canSkip ...bool) int {
	v := getEnv(key, canSkip...)
	i, err := strconv.Atoi(v)
	if err != nil {
		logConvertError(v, key, "int")
	}
	return i
}

func IntList(key string, canSkip ...bool) []int {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]int, len(vs))
	for index, v := range vs {
		i, err := strconv.Atoi(v)
		if err != nil {
			logConvertError(v, key, "int")
		}
		result[index] = i
	}
	return result
}
