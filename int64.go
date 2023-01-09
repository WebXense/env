package env

import (
	"strconv"
	"strings"
)

func Int64(key string, canSkip ...bool) int64 {
	v := getEnv(key, canSkip...)
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		logConvertError(v, key, "int64")
	}
	return i64
}

func Int64List(key string, canSkip ...bool) []int64 {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]int64, len(vs))
	for index, v := range vs {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logConvertError(v, key, "int64")
		}
		result[index] = i64
	}
	return result
}
