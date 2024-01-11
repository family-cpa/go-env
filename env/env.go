package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func init() {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	dir := filepath.Dir(path)
	_ = godotenv.Load(dir + "/.env") // include local .env file
}

// String return env value as string or default value
func String(key string, defaultValue ...string) string {
	return defaultString(os.Getenv(key), defaultValue)
}

// Strings return env value as string slice separated by sep or default value
func Strings(key string, sep string, defaultValue ...[]string) []string {
	value := make([]string, 0)

	for _, item := range strings.Split(String(key), sep) {
		value = append(value, strings.Trim(item, " "))
	}

	if len(value) < 1 && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}

// Int return env value as int or default value
func Int(key string, defaultValue ...int) int {
	value, _ := strconv.Atoi(String(key))

	if value == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}

// Duration return env value as time.Duration parsed from string format or default value
// such as "300ms", "-1.5h" or "2h45m"
// valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
func Duration(key string, defaultValue ...string) time.Duration {
	t, _ := time.ParseDuration(defaultString(os.Getenv(key), defaultValue))
	return t
}

func defaultString(value string, defaultValue []string) string {
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}
