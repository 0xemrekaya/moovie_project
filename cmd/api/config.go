package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

func loadConfig() (*config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		// .env file is optional, so we just log a warning
		return nil, fmt.Errorf(".env file not found, using default values")
	}

	// Build DSN from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "admin")
	dbName := getEnv("DB_NAME", "moovie")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	cfg.db.dsn = dsn

	apiPort := getEnv("API_PORT", "8000")
	apiEnv := getEnv("API_ENV", "development")
	apiDBMaxOpenConns := getEnv("API_DB_MAX_OPEN_CONNS", "25")
	apiDBMaxIdleConns := getEnv("API_DB_MAX_IDLE_CONNS", "25")
	apiDBMaxIdleTime := getEnv("API_DB_MAX_IDLE_TIME", "15m")
	apiLimiterRPS := getEnv("API_LIMITER_RPS", "2")
	apiLimiterBurst := getEnv("API_LIMITER_BURST", "4")
	apiLimiterEnabled := getEnv("API_LIMITER_ENABLED", "true")
	port, err := strconv.Atoi(apiPort)
	if err != nil {
		return nil, fmt.Errorf("API_PORT is not a valid integer")
	}
	cfg.port = port
	cfg.env = apiEnv
	cfg.db.maxOpenConns, err = strconv.Atoi(apiDBMaxOpenConns)
	if err != nil {
		return nil, fmt.Errorf("API_DB_MAX_OPEN_CONNS is not a valid integer")
	}
	cfg.db.maxIdleConns, err = strconv.Atoi(apiDBMaxIdleConns)
	if err != nil {
		return nil, fmt.Errorf("API_DB_MAX_IDLE_CONNS is not a valid integer")
	}
	cfg.db.maxIdleTime, err = time.ParseDuration(apiDBMaxIdleTime)
	if err != nil {
		return nil, fmt.Errorf("API_DB_MAX_IDLE_TIME is not a valid duration")
	}
	cfg.limiter.rps, err = strconv.ParseFloat(apiLimiterRPS, 64)
	if err != nil {
		return nil, fmt.Errorf("API_LIMITER_RPS is not a valid float")
	}
	cfg.limiter.burst, err = strconv.Atoi(apiLimiterBurst)
	if err != nil {
		return nil, fmt.Errorf("API_LIMITER_BURST is not a valid integer")
	}
	cfg.limiter.enabled, err = strconv.ParseBool(apiLimiterEnabled)
	if err != nil {
		return nil, fmt.Errorf("API_LIMITER_ENABLED is not a valid boolean")
	}

	return &cfg, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
