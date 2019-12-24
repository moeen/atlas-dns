package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	sectorId int
	port     int
)

// getIntFromEnv looks for the given key in environment and tries to convert it to integer
func getIntFromEnv(key string) (int, error) {
	s, exists := os.LookupEnv(key)
	if !exists {
		return 0, fmt.Errorf("env variable %s not found", key)
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", key)
	}

	return i, nil
}

// LoadConfig tries to set the needed variables from .env file and then, it will set those to variables defined above
func LoadConfig() error {
	_ = godotenv.Load()

	s, err := getIntFromEnv("DNS_SECTOR_ID")
	if err != nil {
		return err
	}
	sectorId = s

	p, err := getIntFromEnv("DNS_PORT")
	// If DNS_PORT was not found, we will use 8080 as the default port
	if err != nil {
		port = 8080
	} else {
		port = p
	}

	return nil
}

// GetSectorId returns the Sector ID of the application which is loaded from environment variable
func GetSectorId() int {
	return sectorId
}

// GetPort returns the http server port which is either the port from environment variable or 8080
func GetPort() int {
	return port
}
