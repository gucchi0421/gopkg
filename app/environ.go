package app

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}

// skip db load with "make run" or RUN_ENV=test && go run .
func IsTest(b ...bool) bool {
	// if you have a database pass false
	if len(b) > 0 && b[0] == false {
		return false
	}
	return GetEnv("RUN_ENV", "") == "test"
}

// have a db for development
func IsDevelopment(init ...string) bool {
	// please specify "init" before loading godotenv
	if len(init) > 0 && init[0] == "init" {
		return true
	}
	return GetEnv("GO_ENV", "production") == "development"
}

// is deploy stage or GO_ENV not setting
func IsProduction() bool {
	return GetEnv("GO_ENV", "production") == "production"
}
