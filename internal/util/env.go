package util

import "os"

// HasAnyEnv checks if any of the provided environment variable keys exist and are non-empty.
func HasAnyEnv(keys ...string) bool {
	for _, key := range keys {
		if val, exists := os.LookupEnv(key); exists && val != "" {
			return true
		}
	}
	return false
}

// GetEnv fallback utility if you need to fetch specific values later (e.g., WAYLAND_DISPLAY)
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}