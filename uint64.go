package env

import (
	"strconv"
	"strings"
)

func Uint64(key string, canSkip ...bool) uint64 {
	v := getEnv(key, canSkip...)
	u64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		logConvertError(v, key, "uint64")
	}
	return u64
}

func Uint64List(key string, canSkip ...bool) []uint64 {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]uint64, len(vs))
	for index, v := range vs {
		u64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			logConvertError(v, key, "uint64")
		}
		result[index] = u64
	}
	return result
}
