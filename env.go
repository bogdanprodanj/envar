package envar

import (
	"os"
	"strconv"
	"time"
)

// String retrieves string value of the env variable named by the key or returns predefined one
func String(key, def string) string {
	s := os.Getenv(key)
	if len(s) < 1 {
		return def
	}
	return s
}

// Int retrieves int value of the env variable named by the key or returns predefined one.
// Returns an error if value can not be converted to int
func Int(key, def string) (int, error) {
	return strconv.Atoi(String(key, def))
}

// MustInt retrieves int value of the env variable named by the key or returns predefined one.
// Panics if value can not be converted to int
func MustInt(key, def string) int {
	i, err := Int(key, def)
	if err != nil {
		panic(key + " must be int type")
	}
	return i
}

// Float64 retrieves float64 value of the env variable named by the key or returns predefined one.
// Returns an error if value can not be converted to float64
func Float64(key, def string) (float64, error) {
	return strconv.ParseFloat(String(key, def), 64)
}

// MustFloat64 retrieves int value of the env variable named by the key or returns predefined one.
// Panics if value can not be converted to float64
func MustFloat64(key, def string) float64 {
	f, err := Float64(key, def)
	if err != nil {
		panic(key + " must be float type")
	}
	return f
}

// Bool retrieves bool value of the env variable named by the key or returns predefined one.
// Returns an error if value can not be converted to bool
func Bool(key, def string) (bool, error) {
	return strconv.ParseBool(String(key, def))
}

// MustBool retrieves bool value of the env variable named by the key or returns predefined one.
// Panics if value can not be converted to bool
func MustBool(key, fallback string) bool {
	b, err := Bool(key, fallback)
	if err != nil {
		panic(key + " must be boolean type")
	}
	return b
}

// Duration retrieves time.Duration value of the env variable named by the key or returns predefined one.
// Returns an error if value can not be converted to time.Duration
func Duration(key, def string) (time.Duration, error) {
	return time.ParseDuration(String(key, def))
}

// MustDuration retrieves time.Duration value of the env variable named by the key or returns predefined one.
// Panics if value can not be converted to time.Duration
func MustDuration(key, def string) time.Duration {
	d, err := Duration(key, def)
	if err != nil {
		panic(key + " must be time.Duration type (i.e. \"300ms\", \"1.5h\", \"2h45m\") etc.")
	}
	return d
}
