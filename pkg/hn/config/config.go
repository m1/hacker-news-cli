package config

import (
	"os"
	"strconv"
)

type Config struct {
	MaxPosts int
	UrlBase  string
}

func NewConfig() (config *Config, err error) {
	config = &Config{}
	config.MaxPosts, err = strconv.Atoi(getEnv("HN_MAX_POSTS", "100"))
	if err != nil {
		return nil, err
	}

	return config, err
}

// GetEnv returns the env string or the fallback if not found
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}