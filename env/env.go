package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func Get(key string, defaultValue ...string) string {
	return defaultString(os.Getenv(key), defaultValue)
}

func GetInt(key string, defaultValue ...int) int {
	value, err := strconv.Atoi(Get(key))
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
	return value
}

func GetDuration(key string, defaultValue ...string) time.Duration {
	t, _ := time.ParseDuration(defaultString(os.Getenv(key), defaultValue))
	return t
}

func defaultString(value string, defaultValue []string) string {
	if len(value) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}
