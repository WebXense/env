package env

import (
	"strconv"
	"strings"
)

func Float64(key string, canSkip ...bool) float64 {
	v := getEnv(key, canSkip...)
	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		logConvertError(v, key, "float64")
	}
	return f64
}

func Float64List(key string, canSkip ...bool) []float64 {
	vs := strings.Split(getEnv(key, canSkip...), ",")
	result := make([]float64, len(vs))
	for index, v := range vs {
		f64, err := strconv.ParseFloat(v, 64)
		if err != nil {
			logConvertError(v, key, "float64")
		}
		result[index] = f64
	}
	return result
}
