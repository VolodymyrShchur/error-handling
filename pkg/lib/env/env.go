package env

import (
	"os"
	"strconv"
)

// Get returns a value of the environment variable.
func Get(key string) (string, bool) {
	return os.LookupEnv(key)
}

// GetOrDefault returns a value of the environment variable
// or default value if the variable is not set.
func GetOrDefault(key, defval string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defval
	}

	return val
}

// GetInt returns a value of the environment variable converted to int.
func GetInt(key string) (int, bool) {
	val, exists := os.LookupEnv(key)
	if !exists {
		return 0, false
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		return 0, false
	}

	return parsed, true
}

// GetIntOrDefault returns a value of the environment variable converted to int
// or default value if the variable is not set or error appeared while converting to int.
func GetIntOrDefault(key string, defval int) int {
	val, ok := GetInt(key)
	if !ok {
		return defval
	}

	return val
}

// GetBool returns a value of the environment variable converted to bool.
func GetBool(key string) (res, ok bool) {
	val, exists := os.LookupEnv(key)
	if exists {
		parsed, err := strconv.ParseBool(val)

		return parsed, err == nil
	}

	return false, false
}

// GetBoolOrDefault returns a value of the environment variable converted to bool
// or default value if the variable is not set or error appeared while converting to bool.
func GetBoolOrDefault(key string, defval bool) bool {
	val, ok := GetBool(key)
	if !ok {
		return defval
	}

	return val
}
