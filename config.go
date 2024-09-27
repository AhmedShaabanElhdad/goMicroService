package main

import (
	"fmt"
	"os"
)

// better to get this value from the Real time instead of static code
// the whole server environment variable will be inject during the server running
type Config struct {
	port      string
	DBUser    string
	DBPasswd  string
	DBAddr    string
	DBName    string
	JwtSecret string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		port:     getEnv("PORT", ":3000"),
		DBUser:   getEnv("DB_USER", "root"),
		DBPasswd: getEnv("DB_PASSWORD", "password"),
		DBAddr: fmt.Sprintf("%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),
		DBName:    getEnv("DB_NAME", "projectmanager"),
		JwtSecret: getEnv("JWT_SECRET", "randomjswsecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
