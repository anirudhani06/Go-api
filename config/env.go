package config

import "os"

type Config struct {
	DbAddr string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DbAddr: getEnv("DATABASE_URL", "/db/data.db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value

	}
	return fallback
}
