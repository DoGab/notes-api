package utils

import (
	"fmt"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if fallback == "" {
		fmt.Printf("%s is unset and no default value set\n", key)
	}
	return fallback
}
