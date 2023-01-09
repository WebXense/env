package env

import (
	"strconv"
	"strings"
)

func Uint(key string, canSkip ...bool) uint {
	v := getEnv(key, canSkip...)
	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		logConvertError(v, key, "uint")
	}
	return uint(u)
}

func UintList(key string, canSkip ...bool) []uint {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]uint, len(vs))
	for index, v := range vs {
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			logConvertError(v, key, "uint")
		}
		result[index] = uint(u)
	}
	return result
}
