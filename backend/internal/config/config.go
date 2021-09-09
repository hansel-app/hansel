package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Config struct {
	Environment string
	ServerPort  int
	DBHostname  string
	DBPort      int
	DBName      string
	DBUsername  string
	DBPassword  string
}

const (
	envDevelopment        = "development"
	envTest               = "test"
	envProduction         = "production"
	configLoadErrTemplate = "unable to load env vars: %w"
)

var (
	environments        = []string{envDevelopment, envTest, envProduction}
	developmentEnvFiles = []string{".env.development.local", ".env.development"}
	testEnvFiles        = []string{".env.test.local", ".env.test"}
	productionEnvFiles  = []string{".env"}
	projectRootPath     = getProjectRootPath()
)

// Load loads the application configuration from environment variables.
func Load() (*Config, error) {
	var c Config

	err := c.loadEnvironment()
	if err != nil {
		return nil, fmt.Errorf(configLoadErrTemplate, err)
	}

	err = c.loadConfig()
	if err != nil {
		return nil, fmt.Errorf(configLoadErrTemplate, err)
	}

	return &c, nil
}

// loadEnvironment loads the GO_ENV environment variable from the shell environment.
// If GO_ENV is not set, it defaults to the "development" environment.
func (c *Config) loadEnvironment() error {
	env := os.Getenv("GO_ENV")
	if env == "" {
		c.Environment = envDevelopment
		return nil
	}
	for _, environment := range environments {
		if env == environment {
			c.Environment = env
			return nil
		}
	}
	return fmt.Errorf("unrecognised environment: %s", env)
}

func (c *Config) loadConfig() error {
	var configFileNames []string
	switch c.Environment {
	case envDevelopment:
		configFileNames = developmentEnvFiles
	case envTest:
		configFileNames = testEnvFiles
	case envProduction:
		configFileNames = productionEnvFiles
	}
	err := loadEnvFromFiles(configFileNames)
	if err != nil {
		log.Printf("Loading configuration from shell environment...")
	}

	c.ServerPort, err = parsePort(os.Getenv("SERVER_PORT"))
	if err != nil {
		return fmt.Errorf("unable to parse server port: %w", err)
	}
	c.DBHostname = os.Getenv("DB_HOSTNAME")
	c.DBPort, err = parsePort(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("unable to parse db port: %w", err)
	}
	c.DBName = os.Getenv("DB_NAME")
	if c.DBName == "" {
		return errors.New("unable to parse db name: cannot be empty")
	}
	c.DBUsername = os.Getenv("DB_USERNAME")
	c.DBPassword = os.Getenv("DB_PASSWORD")

	return nil
}

// getProjectRootPath returns the root path of the project dynamically. This is in contrast to hard-coded paths
// which break upon building the server into an executable.
func getProjectRootPath() string {
	// Get full path of current file from runtime.
	_, b, _, _ := runtime.Caller(0)

	// Traverse two parent directories up. Note that this is sensitive to the level of nesting
	// of this file within the project structure.
	return filepath.Join(filepath.Dir(b), "../../")
}

// loadEnvFromFile reads in environment variables declared within a file.
func loadEnvFromFile(configFileName string) error {
	err := godotenv.Load(projectRootPath + "/" + configFileName)
	if err != nil {
		return fmt.Errorf("error while loading %s: %w", configFileName, err)
	}
	return nil
}

// loadEnvFromFiles attempts to read in environment variables declared within a list of files sequentially
// until a file is successfully read.
func loadEnvFromFiles(configFileNames []string) error {
	for _, filename := range configFileNames {
		log.Printf("Loading configuration from '%s'...", filename)
		err := loadEnvFromFile(filename)
		// If a config file managed to load, do not load the other files.
		if err == nil {
			return nil
		}
		log.Print(fmt.Sprintf("'%s' not found. Loading next config file...", filename))
	}

	configFileNamesList := strings.Join(configFileNames, ", ")
	return fmt.Errorf("no config file found from any of %s", configFileNamesList)
}

// parsePort converts a string into an integer if it is within the acceptable range for port numbers (0 to 65535).
func parsePort(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("invalid port: not an integer")
	}
	if result < 0 || result > 65535 {
		return 0, errors.New("invalid port: port number should be a 16-bit unsigned integer")
	}
	return result, nil
}
