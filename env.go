package env

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var NotKeyError = errors.New("NotKeyError")

func Load() error {
	return godotenv.Load()
}

func LoadMust() {
	e := godotenv.Load()
	if e != nil {
		panic(e)
	}
}

func Has(key string) bool {
	val := os.Getenv(key)
	if val == "" {
		return false
	}
	return true
}

func Must(key string) {
	if !Has(key) {
		panic(NotKeyError)
	}
}

func BoolDefault(key string, daf bool) bool {
	if strings.ToLower(os.Getenv(key)) == "true" {
		return true
	} else if strings.ToLower(os.Getenv(key)) == "false" {
		return false
	} else {
		return false
	}
}

func BoolMust(key string, daf bool) bool {
	val := os.Getenv(key)
	if val == "" {
		panic(NotKeyError)
	}
	if strings.ToLower(val) == "true" {
		return true
	} else if strings.ToLower(val) == "false" {
		return false
	} else {
		return false
	}
}

func StringDefault(key string, daf string) string {
	val := os.Getenv(key)
	if val == "" {
		return daf
	}
	return val
}

func StringMust(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(NotKeyError)
	}
	return val
}

func IntDefault(key string, daf int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return val
	}
	return daf
}

func IntMust(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(NotKeyError)
	}
	return val
}

func Int64Default(key string, daf int64) int64 {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err == nil {
		return val
	}
	return daf
}

func Int64Must(key string) int64 {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		panic(NotKeyError)
	}
	return val
}

func Float32Default(key string, daf float32) float32 {
	val, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err == nil {
		return float32(val)
	}
	return daf
}

func Float32Must(key string) float32 {
	val, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err != nil {
		panic(NotKeyError)
	}
	return float32(val)
}

func Float64Default(key string, daf float64) float64 {
	val, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err == nil {
		return val
	}
	return daf
}

func Float64Must(key string) float64 {
	val, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		panic(NotKeyError)
	}
	return val
}
