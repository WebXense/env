package env

import (
	"log"
	"os"
)

func getEnv(key string, canSkip ...bool) string {
	v := os.Getenv(key)
	if !parseCanSkip(canSkip...) && v == "" {
		log.Fatalf("env: missing key %s", key)
	}
	return v
}

func parseCanSkip(canSkip ...bool) bool {
	if len(canSkip) == 0 {
		return false
	}
	return canSkip[0]
}

func logConvertError(key, value, convertType string) {
	log.Printf("env: cannot convert %s to %s (key: %s)", value, convertType, key)
}
